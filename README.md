# Anscombe's Quartet Analysis

## Overview

This repository contains implementations of Anscombe's Quartet analysis in multiple programming languages. Our our objective here is to determine if using go is a more suitable approach than using r or python because our company, in this case, wants to identify if go is a faster alternative. We benchmark a statistical linear regression model in order to identify if if go takes longer or less time to compute than r and python. 


- Python (using pandas, statsmodels, matplotlib)
- Go (using the stats package)
- R (LM)

## Requirements

### Python
- Python 3.7+
- pandas
- statsmodels
- matplotlib

### Go
- Go 1.16+
- github.com/montanaflynn/stats package

### R
- R 4.0+
- Base R statistics package

## Installation 
1. Clone this repository
2. Ensure all packages are downloaded correctly

## Using the code

Use the following shell script to run all experiments. Optionally, you may run all experiments separately.

Running all:
```
source run_all.sh
```

# Analysis of Performance

- All datasets showed slopes ≈ 0.5 and intercepts ≈ 3.0
- Results matched Python (using statsmodels) and R (using lm) implementations
- Accuracy was comparable across all three languages

### Performance Comparison
Execution times:

- Go: 0.000174 seconds
- Python: 0.014281 seconds
- R: 0.027000 seconds

## Management Receommendation
As we see the Go program executed way faster than Python and R. We see approximately between an 82 to 155x speedup when using Go compared to R and Python. Reasons for this may include that Go is statically typed and it's a compiled language versus Python which is an interpretive language and R which is mainly used for statistics workflows. This is a clear indicator that Go may be a faster implementation and can serve models at higher inference speeds. In any case this is a perfect use case for Go for production use cases and potentially higher performance computing needs. 

However, despite the massive speedups seen by Go one area of concern is the language's maturity in statistics workflows. Unlike Go, Python has extensive libraries designed specifically for statistics workflows. and often times is the industry standard for model development. Typically, this is used in industry currently and will always have the most support of all three implementations. 

Despite Go's ability to run code incredibly fast, Python still allows the developer to have much more autonomy and free flow in developing their models and running them. Regardless Go does perform faster and there may be some sort of hybrid approach between Python and Go that engineering teams can leverage. Ultimately R is a fantastic language for pure statisticians and also has extensive libraries that practitioners can leverage in order to run statistical analysis. However R is not intended for production and will probably not be the best language to go with when developing production level applications. 