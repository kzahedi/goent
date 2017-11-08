# goent -  GO Implementation of Entropy Measures
[![GoDoc](https://godoc.org/github.com/kzahedi/goent?status.svg)](https://godoc.org/github.com/kzahedi/goent) [![Build Status](https://travis-ci.org/kzahedi/goent.svg?branch=master)](https://travis-ci.org/kzahedi/goent) [![Coverage Status](https://coveralls.io/repos/github/kzahedi/goent/badge.svg?branch=master)](https://coveralls.io/github/kzahedi/goent?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/kzahedi/goent)](https://goreportcard.com/report/github.com/kzahedi/goent)[![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/avelino/awesome-go)

## Measures for discrete state spaces
### Averaged measures
- Entropy
    - Shannon
    - Maximum Likelihood with Bias Correction
    - Horvitz-Thompson 
    - Chao-Shen
- Conditional Entropy
- Mutual Information
- Conditional Mutual Information
- Morphological Computation Measures:
    - MC_W
    - MC_A
    - MC_SY
    - MC_WS
    - MC_MI
    - MC_WA
    - MC_P
- Max Entropy Estimations:
    - Iterative Scaling

### State-dependent measures
- Mutual Information
- Conditional Mutual Information
- Entropy (Shannon)
- Morphological Computation
    - MC_W
    - MC_A
    - MC_WS
    - MC_WA
    - MC_MI



## Measures for continuous state spaces
### Averaged measures
- Kraskov-Stoegbauer-Grassberger, Algorithm 1
- Kraskov-Stoegbauer-Grassberger, Algorithm 2
- Frenzel-Pompe
- Morphological Computation
    - MC_W
    - MC_A
    - MC_WS
    - MC_WA
    - MC_MI


### State-dependent measures 
- Kraskov-Stoegbauer-Grassberger, Algorithm 1
- Kraskov-Stoegbauer-Grassberger, Algorithm 2
- Frenzel-Pompe
- Morphological Computation
    - MC_W
    - MC_A
    - MC_WS
    - MC_WA
    - MC_MI


References:
-  K. Ghazi-Zahedi, C. Langer, and N. Ay. Morphological computation: Synergy of body and brain. Entropy, 19(9), 2017.
- K. Ghazi-Zahedi, R. Deimel, G. Montufar, V. Wall, and O. Brock. Morphological computation: The good, the bad, and the ugly. In IROS 2017, 2017.
-  K. Ghazi-Zahedi, D. F. Haeufle, G. F. Montufar, S. Schmitt, and N. Ay. Evaluating morphological computation in muscle and dc-motor driven models of hopping movements. Frontiers in Robotics and AI, 3(42), 2016.
- K. Ghazi-Zahedi and J. Rauh. Quantifying morphological computation based on an information decomposition of the sensorimotor loop. In Proceedings of the 13th European Conference on Artificial Life (ECAL 2015), pages 70—77, July 2015.
- K. Zahedi and N. Ay. Quantifying morphological computation. Entropy, 15(5):1887–1915, 2013.
- T. M. Cover and J. A. Thomas. Elements of Information Theory, Volume 2nd. Wiley, Hoboken, New Jersey, USA, 2006.
- I. Csiszar. i-divergence geometry of probability distributions and minimization problems. Ann. Probab., 3(1):146–158, 02 1975.
- A. Chao and T.-J. Shen. Nonparametric estimation of shannon’s index of diversity when there are unseen species in sample. Environmental and Ecological Statistics, 10(4):429–443, 2003.
- S. Frenzel and B. Pompe. Partial mutual information for coupling analysis of multivariate time series. Phys. Rev. Lett., 99:204101, Nov 2007.
- A. Kraskov, H. Stoegbauer, and P. Grassberger. Estimating mutual information. Phys. Rev. E, 69:066138, Jun 2004.
