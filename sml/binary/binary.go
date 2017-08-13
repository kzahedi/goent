package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/kzahedi/goent"
	pb "gopkg.in/cheggaaa/pb.v1"
)

func check3DProbabilityDistribution(p [][][]float64) {
	sum := 0.0
	for x := 0; x < len(p); x++ {
		for y := 0; y < len(p[x]); y++ {
			for z := 0; z < len(p[x][y]); z++ {
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

func bin(a int) float64 {
	if a == 0 {
		return -1.0
	}
	return 1.0
}

func pw2_c_w1_a1(w2, w1, a1 int, phi, psi, chi float64) float64 {
	z := math.Exp(phi*bin(w2)*bin(w1) + psi*bin(w2)*bin(a1) + chi*bin(w2)*bin(w1)*bin(a1))
	n := math.Exp(phi*bin(0)*bin(w1)+psi*bin(0)*bin(a1)+chi*bin(0)*bin(w1)*bin(a1)) +
		math.Exp(phi*bin(1)*bin(w1)+psi*bin(1)*bin(a1)+chi*bin(1)*bin(w1)*bin(a1))
	return z / n
}

func pa1_c_s1(a1, s1 int, mu float64) float64 {
	z := math.Exp(mu * bin(a1) * bin(s1))
	n := math.Exp(mu*bin(0)*bin(s1)) + math.Exp(mu*bin(1)*bin(s1))
	return z / n
}

func ps1_c_w1(s1, w1 int, zeta float64) float64 {
	z := math.Exp(zeta * bin(w1) * bin(s1))
	n := math.Exp(zeta*bin(0)*bin(w1)) + math.Exp(zeta*bin(1)*bin(w1))
	return z / n
}

func pw1(w1 int, tau float64) float64 {
	z := math.Exp(tau * bin(w1))
	n := math.Exp(tau*bin(0)) + math.Exp(tau*bin(1))
	return z / n
}

func calculate_MC_W(mu, phi, psi, chi, zeta, tau float64) float64 {

	pw2w1a1 := goent.Create3D(2, 2, 2)

	for w2 := 0; w2 < 2; w2++ {
		for w1 := 0; w1 < 2; w1++ {
			for a1 := 0; a1 < 2; a1++ {
				for s1 := 0; s1 < 2; s1++ {
					pw2w1a1[w2][w1][a1] +=
						pw2_c_w1_a1(w2, w1, a1, phi, psi, chi) *
							pa1_c_s1(a1, s1, mu) *
							ps1_c_w1(s1, w1, zeta) *
							pw1(w1, tau)
				}
			}
		}
	}

	// check3DProbabilityDistribution(pw2w1a1)

	return goent.MC_W(pw2w1a1)
}

func calculate_MC_MI(mu, phi, psi, chi, zeta, tau float64) float64 {

	pw2w1 := goent.Create2D(2, 2)
	pa1s1 := goent.Create2D(2, 2)
	v := 0.0

	for w2 := 0; w2 < 2; w2++ {
		for w1 := 0; w1 < 2; w1++ {
			for a1 := 0; a1 < 2; a1++ {
				for s1 := 0; s1 < 2; s1++ {
					v = pw2_c_w1_a1(w2, w1, a1, phi, psi, chi) *
						pa1_c_s1(a1, s1, mu) *
						ps1_c_w1(s1, w1, zeta) *
						pw1(w1, tau)
					pw2w1[w2][w1] += v
					pa1s1[a1][s1] += v
				}
			}
		}
	}
	return goent.MC_MI(pw2w1, pa1s1)
}

func calculate_MC_A(mu, phi, psi, chi, zeta, tau float64) float64 {

	pw2a1w1 := goent.Create3D(2, 2, 2)

	for w2 := 0; w2 < 2; w2++ {
		for w1 := 0; w1 < 2; w1++ {
			for a1 := 0; a1 < 2; a1++ {
				for s1 := 0; s1 < 2; s1++ {
					pw2a1w1[w2][a1][w1] +=
						pw2_c_w1_a1(w2, w1, a1, phi, psi, chi) *
							pa1_c_s1(a1, s1, mu) *
							ps1_c_w1(s1, w1, zeta) *
							pw1(w1, tau)
				}
			}
		}
	}

	return goent.MC_A(pw2a1w1)
}

func main() {

	muStr := flag.String("mu", "0", "s -> a. can take list (1,2,3) or range with delta (0:0.1:1.0)")
	phiStr := flag.String("phi", "0:0.1:5", "w -> w'. Can take list (1,2,3) or range with delta (0:0.1:1.0)")
	psiStr := flag.String("psi", "0:0.1:5", "(a -> w'. Can take list (1,2,3) or range with delta (0:0.1:1.0)")
	chiStr := flag.String("chi", "0:0.1:5", "a,w -> w'. Can take list (1,2,3) or range with delta (0:0.1:1.0)")
	zetaStr := flag.String("zeta", "0", "w -> s. Can take list (1,2,3) or range with delta (0:0.1:1.0)")
	tauStr := flag.String("tau", "0", "p(w). Can take list (1,2,3) or range with delta (0:0.1:1.0)")
	mc := flag.String("mc", "MC_W", "quantification to use: MC_W, MC_A, MC_MI (soon: MC_SY, MC_SY_NIS, MC_SY_GIS, MC_SY_SCGIS)")
	verbose := flag.Bool("v", false, "verbose")
	output := flag.String("o", "out.csv", "output file. Default: out.csv")

	flag.Parse()

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
		fmt.Println("out:", *output)
	}

	f, err := os.Create(*output)
	check(err)
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	s := []string{"mu", "phi", "psi", "zeta", "tau", "r"}
	err = writer.Write(s)
	check(err)

	r := 0.0

	bar := pb.StartNew(len(mu) * len(phi) * len(psi) * len(chi) * len(zeta) * len(tau))
	for _, vmu := range mu {
		for _, vphi := range phi {
			for _, vpsi := range psi {
				for _, vchi := range chi {
					for _, vzeta := range zeta {
						for _, vtau := range tau {
							bar.Increment()
							switch *mc {
							case "MC_W":
								r = calculate_MC_W(vmu, vphi, vpsi, vchi, vzeta, vtau)
								s = []string{f2s(vmu), f2s(vphi), f2s(vpsi), f2s(vzeta), f2s(vtau), f2s(r)}
								err = writer.Write(s)
								check(err)
							case "MC_A":
								r = calculate_MC_A(vmu, vphi, vpsi, vchi, vzeta, vtau)
								s = []string{f2s(vmu), f2s(vphi), f2s(vpsi), f2s(vzeta), f2s(vtau), f2s(r)}
								err = writer.Write(s)
								check(err)
							case "MC_MI":
								r = calculate_MC_MI(vmu, vphi, vpsi, vchi, vzeta, vtau)
								s = []string{f2s(vmu), f2s(vphi), f2s(vpsi), f2s(vzeta), f2s(vtau), f2s(r)}
								err = writer.Write(s)
								check(err)
							default:
								panic(fmt.Sprintf("Unknown quantification given %s", *mc))
							}
						}
					}
				}
			}
		}
	}
	bar.FinishPrint("Finished")
}
