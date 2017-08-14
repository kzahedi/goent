package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"
	"strings"
	"sync"

	"github.com/kzahedi/goent"
	pb "gopkg.in/cheggaaa/pb.v1"
)

var mutex sync.Mutex
var mc_sy_iterations int
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write mem profile to file")

type Indicator struct {
	a      int
	b      int
	c      int
	labmda float64
}

func lambda2(ind Indicator, aa, bb int) float64 {
	if ind.a == aa && ind.b == bb {
		return ind.labmda
	}
	return -ind.labmda
}

func lambda3(ind Indicator, aa, bb, cc int) float64 {
	if ind.a == aa && ind.b == bb && ind.c == cc {
		return ind.labmda
	}
	return -ind.labmda
}

func vmap(v int, bins int) float64 {
	return 2.0*float64(v)/float64(bins-1) - 1.0
}

func get_value2(a, b int, factor float64, bins int) float64 {
	return factor * vmap(a, bins) * vmap(b, bins)
}

func get_value3(a, b, c int, factor float64, bins int) float64 {
	return factor * vmap(a, bins) * vmap(b, bins) * vmap(c, bins)
}

func (ind *Indicator) String() string {
	return fmt.Sprintf("Indicator %d %d = %f", ind.a, ind.b, ind.labmda)
}

func check3DProbabilityDistribution(p [][][]float64) {
	sum := 0.0
	for x := 0; x < len(p); x++ {
		for y := 0; y < len(p[x]); y++ {
			for z := 0; z < len(p[x][y]); z++ {
				if math.IsNaN(p[x][y][z]) {
					panic("NaN")
				}
				sum += p[x][y][z]
			}
		}
	}
	if math.Abs(sum-1.0) > 0.0000001 {
		panic(fmt.Sprintf("Does not sum up to one %f", sum))
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func f2s(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}

func getvalues(str string) []float64 {
	var r []float64
	values := strings.Split(str, ":")

	start, err := strconv.ParseFloat(values[0], 64)
	check(err)

	end := start
	delta := start - end + 1.0

	if len(values) == 3 {
		delta, err = strconv.ParseFloat(values[1], 64)
		check(err)

		end, err = strconv.ParseFloat(values[2], 64)
		check(err)
	}

	for v := start; v <= end; v += delta {
		r = append(r, v)
	}

	return r
}

func pw2_c_w1_a1(w2, w1, a1, bins int,
	phi, psi, chi float64,
	w2w1a1i, w2w1i, w2a1i []Indicator) float64 {
	z := 0.0
	n := 0.0
	for _, ind := range w2w1a1i {
		z += lambda3(ind, w2, w1, a1)
	}
	for _, ind := range w2w1i {
		z += lambda2(ind, w2, w1)
	}
	for _, ind := range w2a1i {
		z += lambda2(ind, w2, a1)
	}
	z = math.Exp(z)

	for w22 := 0; w22 < bins; w22++ {
		nn := 0.0

		for _, ind := range w2w1a1i {
			nn += lambda3(ind, w22, w1, a1)
		}
		for _, ind := range w2w1i {
			nn += lambda2(ind, w22, w1)
		}
		for _, ind := range w2a1i {
			nn += lambda2(ind, w22, a1)
		}
		n += math.Exp(nn)
	}

	return z / n
}

func pa1_c_s1(a1, s1, bins int, mu float64, a1s1i []Indicator) float64 {
	z := 0.0
	n := 0.0
	nn := 0.0
	for _, ind := range a1s1i {
		z += lambda2(ind, a1, s1)
	}
	z = math.Exp(z)

	for a11 := 0; a11 < bins; a11++ {
		nn = 0.0
		for _, ind := range a1s1i {
			nn += lambda2(ind, a11, s1)
		}
		n += math.Exp(nn)
	}
	return z / n
}

func ps1_c_w1(s1, w1, bins int, zeta float64, s1w1i []Indicator) float64 {
	z := 0.0
	n := 0.0
	nn := 0.0

	for _, ind := range s1w1i {
		z += lambda2(ind, s1, w1)
	}
	z = math.Exp(z)

	for s11 := 0; s11 < bins; s11++ {
		nn = 0.0
		for _, ind := range s1w1i {
			nn += lambda2(ind, s11, w1)
		}
		n += math.Exp(nn)
	}

	return z / n
}

func pw1(w1, bins int, tau float64) float64 {
	return 1.0 / float64(bins)
}

func generate_w2w1a1_indicators(chi float64, bins int) []Indicator {
	var r []Indicator
	for w2 := 0; w2 < bins; w2++ {
		for w1 := 0; w1 < bins; w1++ {
			for a1 := 0; a1 < bins; a1++ {
				r = append(r, Indicator{w2, w1, a1, get_value3(w2, w1, a1, chi, bins)})
			}
		}
	}
	// fmt.Println("w2w1a1")
	// fmt.Println(r)
	return r
}

func generate_w2w1_indicators(phi float64, bins int) []Indicator {
	var r []Indicator
	for w2 := 0; w2 < bins; w2++ {
		for w1 := 0; w1 < bins; w1++ {
			r = append(r, Indicator{w2, w1, -1, get_value2(w2, w1, phi, bins)})
		}
	}
	return r
}

func generate_w2a1_indicators(psi float64, bins int) []Indicator {
	var r []Indicator
	for w2 := 0; w2 < bins; w2++ {
		for a1 := 0; a1 < bins; a1++ {
			r = append(r, Indicator{w2, a1, -1, get_value2(w2, a1, psi, bins)})
		}
	}
	return r
}

func generate_a1s1_indicators(mu float64, bins int) []Indicator {
	var r []Indicator
	for a1 := 0; a1 < bins; a1++ {
		for s1 := 0; s1 < bins; s1++ {
			r = append(r, Indicator{a1, s1, -1, get_value2(a1, s1, mu, bins)})
		}
	}
	return r
}

func generate_s1w1_indicators(zeta float64, bins int) []Indicator {
	var r []Indicator
	for s1 := 0; s1 < bins; s1++ {
		for w1 := 0; w1 < bins; w1++ {
			r = append(r, Indicator{s1, w1, -1, get_value2(s1, w1, zeta, bins)})
		}
	}
	return r
}

func write(writer *os.File, phi, psi, chi, mu, zeta, tau, r float64) {
	mutex.Lock()
	defer mutex.Unlock()
	s := fmt.Sprintf("%f,%f,%f,%f,%f,%f,%f\n", phi, psi, chi, mu, zeta, tau, r)
	writer.WriteString(s)
	writer.Sync()
}

func calculate_MC_W(mu, phi, psi, chi, zeta, tau float64, bins int, writer *os.File) {

	pw2w1a1 := goent.Create3D(2, 2, 2)
	w2w1a1i := generate_w2w1a1_indicators(chi, bins)
	w2w1i := generate_w2w1_indicators(phi, bins)
	w2a1i := generate_w2a1_indicators(psi, bins)
	a1s1i := generate_a1s1_indicators(mu, bins)
	s1w1i := generate_s1w1_indicators(zeta, bins)

	for w2 := 0; w2 < 2; w2++ {
		for w1 := 0; w1 < 2; w1++ {
			for a1 := 0; a1 < 2; a1++ {
				for s1 := 0; s1 < 2; s1++ {
					pw2w1a1[w2][w1][a1] +=
						pw2_c_w1_a1(w2, w1, a1, bins, phi, psi, chi, w2w1a1i, w2w1i, w2a1i) *
							pa1_c_s1(a1, s1, bins, mu, a1s1i) *
							ps1_c_w1(s1, w1, bins, zeta, s1w1i) *
							pw1(w1, bins, tau)
				}
			}
		}
	}

	r := goent.MC_W(pw2w1a1)
	write(writer, phi, psi, chi, mu, zeta, tau, r)
}

func calculate_MC_MI(mu, phi, psi, chi, zeta, tau float64, bins int, writer *os.File) {

	pw2w1 := goent.Create2D(2, 2)
	pa1s1 := goent.Create2D(2, 2)
	w2w1a1i := generate_w2w1a1_indicators(chi, bins)
	w2w1i := generate_w2w1_indicators(phi, bins)
	w2a1i := generate_w2a1_indicators(psi, bins)
	a1s1i := generate_a1s1_indicators(mu, bins)
	s1w1i := generate_s1w1_indicators(zeta, bins)

	v := 0.0

	for w2 := 0; w2 < 2; w2++ {
		for w1 := 0; w1 < 2; w1++ {
			for a1 := 0; a1 < 2; a1++ {
				for s1 := 0; s1 < 2; s1++ {
					v = pw2_c_w1_a1(w2, w1, a1, bins, phi, psi, chi, w2w1a1i, w2w1i, w2a1i) *
						pa1_c_s1(a1, s1, bins, mu, a1s1i) *
						ps1_c_w1(s1, w1, bins, zeta, s1w1i) *
						pw1(w1, bins, tau)
					pw2w1[w2][w1] += v
					pa1s1[a1][s1] += v
				}
			}
		}
	}

	r := goent.MC_MI(pw2w1, pa1s1)
	write(writer, phi, psi, chi, mu, zeta, tau, r)
	// s := []string{f2s(phi), f2s(psi), f2s(chi), f2s(mu), f2s(zeta), f2s(tau), f2s(r)}
	// check(writer.Write(s))
}

func calculate_MC_A(mu, phi, psi, chi, zeta, tau float64, bins int, writer *os.File) {

	pw2w1a1 := goent.Create3D(2, 2, 2)
	w2w1a1i := generate_w2w1a1_indicators(chi, bins)
	w2w1i := generate_w2w1_indicators(phi, bins)
	w2a1i := generate_w2a1_indicators(psi, bins)
	a1s1i := generate_a1s1_indicators(mu, bins)
	s1w1i := generate_s1w1_indicators(zeta, bins)

	for w2 := 0; w2 < 2; w2++ {
		for w1 := 0; w1 < 2; w1++ {
			for a1 := 0; a1 < 2; a1++ {
				for s1 := 0; s1 < 2; s1++ {
					pw2w1a1[w2][w1][a1] +=
						pw2_c_w1_a1(w2, w1, a1, bins, phi, psi, chi, w2w1a1i, w2w1i, w2a1i) *
							pa1_c_s1(a1, s1, bins, mu, a1s1i) *
							ps1_c_w1(s1, w1, bins, zeta, s1w1i) *
							pw1(w1, bins, tau)
				}
			}
		}
	}

	r := goent.MC_A(pw2w1a1)
	// s := []string{f2s(phi), f2s(psi), f2s(chi), f2s(mu), f2s(zeta), f2s(tau), f2s(r)}
	// check(writer.Write(s))

	write(writer, phi, psi, chi, mu, zeta, tau, r)
}

func calculate_MC_SY(mu, phi, psi, chi, zeta, tau float64, bins int, writer *os.File) {

	pw2w1a1 := goent.Create3D(2, 2, 2)
	w2w1a1i := generate_w2w1a1_indicators(chi, bins)
	w2w1i := generate_w2w1_indicators(phi, bins)
	w2a1i := generate_w2a1_indicators(psi, bins)
	a1s1i := generate_a1s1_indicators(mu, bins)
	s1w1i := generate_s1w1_indicators(zeta, bins)

	for w2 := 0; w2 < 2; w2++ {
		for w1 := 0; w1 < 2; w1++ {
			for a1 := 0; a1 < 2; a1++ {
				for s1 := 0; s1 < 2; s1++ {
					pw2w1a1[w2][w1][a1] +=
						pw2_c_w1_a1(w2, w1, a1, bins, phi, psi, chi, w2w1a1i, w2w1i, w2a1i) *
							pa1_c_s1(a1, s1, bins, mu, a1s1i) *
							ps1_c_w1(s1, w1, bins, zeta, s1w1i) *
							pw1(w1, bins, tau)
				}
			}
		}
	}

	r := goent.MC_SY(pw2w1a1, mc_sy_iterations)
	// s := []string{f2s(phi), f2s(psi), f2s(chi), f2s(mu), f2s(zeta), f2s(tau), f2s(r)}
	// check(writer.Write(s))
	write(writer, phi, psi, chi, mu, zeta, tau, r)
}

func main() {

	muStr := flag.String("mu", "0", "s -> a. can take list (1,2,3) or range with delta (0:0.1:1.0)")
	phiStr := flag.String("phi", "0:0.01:5", "w -> w'. Can take list (1,2,3) or range with delta (0:0.1:1.0)")
	psiStr := flag.String("psi", "0:0.01:5", "(a -> w'. Can take list (1,2,3) or range with delta (0:0.1:1.0)")
	chiStr := flag.String("chi", "0:0.1:5", "a,w -> w'. Can take list (1,2,3) or range with delta (0:0.1:1.0)")
	zetaStr := flag.String("zeta", "0", "w -> s. Can take list (1,2,3) or range with delta (0:0.1:1.0)")
	tauStr := flag.String("tau", "0", "p(w). Can take list (1,2,3) or range with delta (0:0.1:1.0)")
	mc := flag.String("mc", "MC_W", "quantification to use: MC_W, MC_A, MC_MI (soon: MC_SY, MC_SY_NIS, MC_SY_GIS, MC_SY_SCGIS)")
	verbose := flag.Bool("v", false, "verbose")
	bins := flag.Int("b", 2, "Bins")
	syci := flag.Int("syci", 500, "MC_SY convergence iterations")
	cpus := flag.Int("cpus", 0, "Nr. of CPUs")
	output := flag.String("o", "out.csv", "output file. Default: out.csv")

	mc_sy_iterations = *syci

	flag.Parse()

	if *cpuprofile != "" {
		fc, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(fc)
		defer pprof.StopCPUProfile()
	}

	if *cpus > 0 {
		runtime.GOMAXPROCS(*cpus)
	} else {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	mu := getvalues(*muStr)
	phi := getvalues(*phiStr)
	psi := getvalues(*psiStr)
	chi := getvalues(*chiStr)
	zeta := getvalues(*zetaStr)
	tau := getvalues(*tauStr)

	if *verbose == true {
		fmt.Println("mu:", *muStr)
		fmt.Println("mu:", mu)
		fmt.Println("phi:", *phiStr)
		fmt.Println("phi:", phi)
		fmt.Println("psi:", *psiStr)
		fmt.Println("psi:", psi)
		fmt.Println("chi:", *chiStr)
		fmt.Println("chi:", chi)
		fmt.Println("zeta:", *zetaStr)
		fmt.Println("zeta:", zeta)
		fmt.Println("tau:", *tauStr)
		fmt.Println("tau:", tau)
		fmt.Println("mc:", *mc)
		fmt.Println("bins:", *bins)
		fmt.Println("out:", *output)
	}

	f, err := os.Create(*output)
	check(err)
	defer f.Close()

	// s := []string{"phi", "psi", "chi", "mu", "zeta", "tau", "MC"}
	f.WriteString(fmt.Sprintf("# phi, psi, chi, mu, zeta, tau, %s\n", *mc))

	var wg sync.WaitGroup

	var mcFunc func(float64, float64, float64, float64, float64, float64, int, *os.File)

	switch *mc {
	case "MC_W":
		mcFunc = calculate_MC_W
	case "MC_A":
		mcFunc = calculate_MC_A
	case "MC_MI":
		mcFunc = calculate_MC_MI
	case "MC_SY":
		mcFunc = calculate_MC_SY
	default:
		panic(fmt.Sprintf("Unknown quantification given %s", *mc))
	}

	bar := pb.StartNew(len(mu) * len(phi) * len(psi) * len(chi) * len(zeta) * len(tau))

	// uiprogress.Start()            // start rendering
	// bar := uiprogress.AddBar(len(mu) * len(phi) * len(psi) * len(chi) * len(zeta) * len(tau)))
	// bar.AppendCompleted()
	// bar.PrependElapsed()

	for _, vmu := range mu {
		for _, vphi := range phi {
			for _, vpsi := range psi {
				for _, vchi := range chi {
					for _, vzeta := range zeta {
						for _, vtau := range tau {
							wg.Add(1)
							go func(vvphi, vvpsi, vvchi, vvmu, vvzeta, vvtau float64, bbins int, ff *os.File) {
								// fmt.Println(fmt.Sprintf("calling mcFunc with %f,%f,%f,%f,%f,%f", vvphi, vvpsi, vvchi, vvmu, vvzeta, vvtau))
								mcFunc(vvphi, vvpsi, vvchi, vvmu, vvzeta, vvtau, bbins, ff)
								wg.Done()
								bar.Increment()
							}(vmu, vphi, vpsi, vchi, vzeta, vtau, *bins, f)
						}
					}
				}
			}
		}
	}
	wg.Wait()
	bar.FinishPrint("Finished")
	if *memprofile != "" {
		fm, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(fm)
		fm.Close()
		return
	}

}
