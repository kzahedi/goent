# goent -  GO Implementation of Entropy Measures
[![GoDoc](https://godoc.org/github.com/kzahedi/goent?status.svg)](https://godoc.org/github.com/kzahedi/goent) [![Build Status](https://travis-ci.org/kzahedi/goent.svg?branch=master)](https://travis-ci.org/kzahedi/goent) [![Coverage Status](https://coveralls.io/repos/github/kzahedi/goent/badge.svg?branch=master)](https://coveralls.io/github/kzahedi/goent?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/kzahedi/goent)](https://goreportcard.com/report/github.com/kzahedi/goent) [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/avelino/awesome-go)

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
- Max Entropy Estimations:
    - Iterative Scaling
- Information Decomposition (Bertschinger et al., 2014) for binary variables

- Morphological Computation measures have been moved to gomi
  https://github.com/kzahedi/gomi

### State-dependent measures
- Mutual Information
- Conditional Mutual Information
- Entropy (Shannon)

- Morphological Computation measures have been moved to gomi
  https://github.com/kzahedi/gomi


## Measures for continuous state spaces
### Averaged measures
- Kraskov-Stoegbauer-Grassberger, Algorithm 1
- Kraskov-Stoegbauer-Grassberger, Algorithm 2
- Frenzel-Pompe

- Morphological Computation measures have been moved to gomi
  https://github.com/kzahedi/gomi


### State-dependent measures
- Kraskov-Stoegbauer-Grassberger, Algorithm 1
- Kraskov-Stoegbauer-Grassberger, Algorithm 2
- Frenzel-Pompe

- Morphological Computation measures have been moved to gomi
  https://github.com/kzahedi/gomi


References:
- T. M. Cover and J. A. Thomas. Elements of Information Theory, Volume 2nd. Wiley, Hoboken, New Jersey, USA, 2006.
- I. Csiszar. i-divergence geometry of probability distributions and minimization problems. Ann. Probab., 3(1):146–158, 02 1975.
- A. Chao and T.-J. Shen. Nonparametric estimation of shannon’s index of diversity when there are unseen species in sample. Environmental and Ecological Statistics, 10(4):429–443, 2003.
- S. Frenzel and B. Pompe. Partial mutual information for coupling analysis of multivariate time series. Phys. Rev. Lett., 99:204101, Nov 2007.
- A. Kraskov, H. Stoegbauer, and P. Grassberger. Estimating mutual information. Phys. Rev. E, 69:066138, Jun 2004.
Bertschinger2013aQuantifying
- N. Bertschinger, J. Rauh, E. Olbrich, J. Jost, and N. Ay, Quantifying unique information, CoRR, 2013

