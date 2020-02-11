# Estimation of R0

#### Author: Yijun Wang
#### Date: 2020-02-09

### To Replicate the Result:
- Download my Jupyter notebook file: [Estimation of R0.ipynb](https://github.com/yijunwang0805/YijunWang/blob/master/Estimation%20of%20R0_Yijun/Estimation%20of%20R0.ipynb)

### Background
> 1. R0 is known as basic reproduction number. It can be understood as the expected number of cases infected by one case in a population where all individuals are susceptible to infection. If R0 < 1, an epidemic will not start. If R0 > 1, an epidemic will be able to start spreading out. If R0 = 1, an epidemic will become an endemic.
> 2. Generation period, or known as serial interval, is assumed to be 7.5, taken from reference 1 and 2
> 3. P value is 0.695, taken from People's Daily Weibo. https://m.weibo.cn/u/2803301701
> 4. R0 estimation formula refers to reference 3. I will upload a proof of formula in a couple of days.

The basic reproduction number R0 has important implication. The greater the R0, the harder to control an epidemic. R0 for SARS and Ebola virus is 0.49 and 1.51, respectively (Gerardo 2004; Althaus, 2014). 

Tracing R0 for nCoV throughout time will provide a glance of the change in R0, giving clues of the effect of social and non-pharmaceutical prevention.

### Reference
1. Li, Q., Guan, X., et al. (2020, January 29). Early Transmission Dynamics in Wuhan, China, of Novel Coronavirus–Infected Pneumonia. The New England Journal of Medicine. https://www.nejm.org/doi/full/10.1056/NEJMoa2001316#article_references

2. Wu, Joseph T., et al. (2020, January 28). Nowcasting and Forecasting the potential domestic and International Spread of the 2019-nCoV Outbreak Originating in Wuhan, China: a Modeling Study. Lancet.https://www.thelancet.com/journals/lancet/article/PIIS0140-6736(20)30260-9/fulltext

3. Zhou Tao, Liu, Y., et al. (2020, January 29). Preliminary Prediction of the Basic Repreduction Number of the Novel Coronavirus 2019-nCoV http://kns.cnki.net/kcms/detail/51.1656.r.20200204.1640.002.html

4. Althaus, Christian L. (2014, June). Estimating the Reproduction Number of Ebola Virus (EBOV) During the 2014 Outbreak in West Africa. PLoS Currents. PMC 4169395. PMID 25642364.

5. Gerardo Chowell, Carlos, P., et al. (2004, July). Model Parameters and Outbreak Control for SARS.
