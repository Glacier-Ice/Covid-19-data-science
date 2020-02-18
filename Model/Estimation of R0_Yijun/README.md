# Estimation of R<sub>0</sub>

#### Author: Yijun Wang
#### Date: 2020-02-09
#### Last Update: 2020-02-16

### Usage
- Download my Jupyter notebook file: [Estimation of R0.ipynb](https://github.com/yijunwang0805/YijunWang/blob/master/Estimation%20of%20R0_Yijun/Estimation%20of%20R0.ipynb)
- The code will load data from a API connection, which is provided by [BlankerL](https://lab.isaaclin.cn/nCoV/). 
- ```R0Func()``` is the function that calculates the basic reproduction number. Its ```inputs``` are the number of ```confirm``` cases, the number of ```suspect``` cases, and days ```t``` since the start of the epidemic. Here, we use the December 1st, 2019 as the start of the epidemic, which is the first nCoV case reported. 

### Summary
- This study seeks to estimate the basic reproduction number by deriving R<sub>0</sub> from the SEIR model. As of 2020-02-14, R<sub>0</sub> is estimated to be 2.41.

### Background
- There has been a novel coronavirus (2019-nCoV) pneumonia outbreak in China since December 2019 which spreads internationally. This study seeks to quantify the basic reproduction number of 2019-nCoV throughout the outbreak.
- R<sub>0</sub> is known as basic reproduction number. It can be understood as the expected number of cases infected by one case in a population where all individuals are susceptible to infection. If R<sub>0</sub> < 1, an epidemic will not start. If R<sub>0</sub> > 1, an epidemic will be able to start spreading out. If R<sub>0</sub> = 1, an epidemic will become an endemic.
- The basic reproduction number R<sub>0</sub> has important implication. The greater the R<sub>0</sub>, the harder to control an epidemic. R<sub>0</sub> for SARS and Ebola virus is 0.49 and 1.51, respectively (Gerardo 2004; Althaus, 2014). 
- Tracing R<sub>0</sub> for nCoV throughout time will provide a glance of the change in R<sub>0</sub>, giving clues of the effectiveness of social and non-pharmaceutical prevention.
- On 2019-12-01, the first case of nCoV exhibits symptons, according to [Huang's Clinical features of patients infected with 2019 novel coronavirus in Wuhan, China](https://www.thelancet.com/journals/lancet/article/PIIS0140-6736(20)30183-5/fulltext#seccestitle10). This date is used as the first day of the epidemic.
- On 2020-1-23, Wuhan city lockdowns [(ifeng news)](http://news.ifeng.com/c/7tpL47zV2Vy). Before Wuhan lockdown, 5 million people leave the city [(Tencent News)](https://new.qq.com/sv1/qd/aoyou.html?cmsid=20200127A0EFXJ00)
- There is a time lag between onset of symptons and case confirmed, particularly due to staff and medical supply shortage
- After the Wuhan outbreak, people quaratine at home if their cases are not severe, which is a common practice to avoid cross-infection. 

### Assumption
- This model assume that infected individuals were not infectious during the incubation period (Zhou, 2020).
- There is no growth or dead during the epidemic.

### Model
A typical **SEIR** (susceptible, exposed, infectious, removed) model can be described as a system of differential equations

<a href="https://www.codecogs.com/eqnedit.php?latex=\small&space;\frac{dS}{dt}&space;=&space;-\beta&space;\frac{S(t)I(t)}{N}&space;\newline&space;\frac{dE}{dt}&space;=&space;\beta&space;\frac{S(t)I(t)}{N}&space;-&space;\alpha&space;E(t)&space;\newline&space;\frac{dI}{dt}&space;=&space;\alpha&space;E(t)&space;-&space;\gamma&space;I(t)&space;\newline&space;\frac{dR(t)}{dt}&space;=&space;\gamma&space;I(t)" target="_blank"><img src="https://latex.codecogs.com/gif.latex?\small&space;\frac{dS}{dt}&space;=&space;-\beta&space;\frac{S(t)I(t)}{N}&space;\newline&space;\frac{dE}{dt}&space;=&space;\beta&space;\frac{S(t)I(t)}{N}&space;-&space;\alpha&space;E(t)&space;\newline&space;\frac{dI}{dt}&space;=&space;\alpha&space;E(t)&space;-&space;\gamma&space;I(t)&space;\newline&space;\frac{dR(t)}{dt}&space;=&space;\gamma&space;I(t)" title="\small \frac{dS}{dt} = -\beta \frac{S(t)I(t)}{N} \newline \frac{dE}{dt} = \beta \frac{S(t)I(t)}{N} - \alpha E(t) \newline \frac{dI}{dt} = \alpha E(t) - \gamma I(t) \newline \frac{dR(t)}{dt} = \gamma I(t)" /></a>

where,

S(t) is the number of susceptible at time t

E(t) is the number of exposed at time t

I(t) is the number of infectious at time t

R(t) is the number of removed, which includes the number of recovered and dead at time t

N(t) is the population at time t

N(t) = S(t) + E(t) + I(t) + R(t)

At the beginning, the ratio of infectious to population is almost negligible.

Therefore,

<a href="https://www.codecogs.com/eqnedit.php?latex=\small&space;t&space;\to&space;0,&space;S(0)&space;\to&space;N(0)" target="_blank"><img src="https://latex.codecogs.com/gif.latex?\small&space;t&space;\to&space;0,&space;S(0)&space;\to&space;N(0)" title="\small t \to 0, S(0) \to N(0)" /></a>

R<sub>0</sub> can be expressed as

<a href="https://www.codecogs.com/eqnedit.php?latex=R_0&space;=&space;(1&plus;\frac{\lambda}{\alpha})(1&plus;\frac{\lambda}{\gamma})" target="_blank"><img src="https://latex.codecogs.com/gif.latex?R_0&space;=&space;(1&plus;\frac{\lambda}{\alpha})(1&plus;\frac{\lambda}{\gamma})" title="R_0 = (1+\frac{\lambda}{\alpha})(1+\frac{\lambda}{\gamma})" /></a>

Given that 

<a href="https://www.codecogs.com/eqnedit.php?latex=\small&space;T_L&space;=&space;\frac{1}{\alpha}&space;\newline&space;T_I&space;=&space;\frac{1}{\gamma}&space;\newline&space;T_G&space;=&space;T_L&space;&plus;&space;T_I" target="_blank"><img src="https://latex.codecogs.com/gif.latex?\small&space;T_L&space;=&space;\frac{1}{\alpha}&space;\newline&space;T_I&space;=&space;\frac{1}{\gamma}&space;\newline&space;T_G&space;=&space;T_L&space;&plus;&space;T_I" title="\small T_L = \frac{1}{\alpha} \newline T_I = \frac{1}{\gamma} \newline T_G = T_L + T_I" /></a>

where,

T<sub>L</sub> is the time of latent period

T<sub>I</sub> is the time of infectious period

T<sub>G</sub> is the time of generation period, which is also known as the serial interval

R<sub>0</sub> can be derived as

<a href="https://www.codecogs.com/eqnedit.php?latex=R_0&space;=&space;1&space;&plus;&space;\lambda&space;T_G&space;&plus;&space;\rho(1&space;-&space;\rho)&space;(\lambda&space;T_G)^2" target="_blank"><img src="https://latex.codecogs.com/gif.latex?R_0&space;=&space;1&space;&plus;&space;\lambda&space;T_G&space;&plus;&space;\rho(1&space;-&space;\rho)&space;(\lambda&space;T_G)^2" title="R_0 = 1 + \lambda T_G + \rho(1 - \rho) (\lambda T_G)^2" /></a>

This will be the formula estimating R<sub>0</sub> in the code. Here is the **mathematical proof** for the equation above

<a href="https://www.codecogs.com/eqnedit.php?latex=R_0&space;=&space;(1&plus;\frac{\lambda}{\alpha})(1&plus;\frac{\lambda}{\gamma})&space;\newline&space;R_0&space;=&space;(1&plus;\lambda&space;T_L)(1&plus;\lambda&space;T_I)&space;\newline&space;R_0&space;=&space;1&space;&plus;&space;\lambda&space;T_I&space;&plus;&space;\lambda&space;T_L&space;&plus;&space;\lambda^2&space;T_I&space;T_L&space;\newline&space;R_0&space;=&space;1&space;&plus;&space;\lambda&space;(T_I&space;&plus;&space;T_L)&space;&plus;&space;\lambda^2&space;T_I&space;T_G&space;\frac{T_L}{T_G}&space;\newline&space;R_0&space;=&space;1&space;&plus;&space;\lambda&space;T_G&space;&plus;&space;\lambda^2&space;T_I&space;T_G&space;\rho&space;\newline&space;R_0&space;=&space;1&space;&plus;&space;\lambda&space;T_G&space;&plus;&space;\lambda^2&space;\rho(T_G-T_L)T_G&space;\newline&space;R_0&space;=&space;1&space;&plus;&space;\lambda&space;T_G&space;&plus;&space;\lambda^2&space;\rho&space;T_G^2&space;-&space;\lambda^2\rho&space;T_G&space;T_L&space;\newline&space;R_0&space;=&space;1&space;&plus;&space;\lambda&space;T_G&space;&plus;&space;\lambda^2&space;\rho&space;T_G^2&space;-&space;\lambda^2\rho&space;T_G&space;\frac{T_L}{T_G}T_G&space;\newline&space;R_0&space;=&space;1&space;&plus;&space;\lambda&space;T_G&space;&plus;&space;\lambda^2&space;\rho&space;T_G^2&space;-&space;\lambda^2\rho&space;T_G&space;\rho&space;T_G&space;\newline&space;R_0&space;=&space;1&space;&plus;&space;\lambda&space;T_G&space;&plus;&space;\lambda^2&space;\rho&space;T_G^2&space;-&space;\lambda^2&space;\rho^2&space;T_G^2&space;\newline&space;R_0&space;=&space;1&space;&plus;&space;\lambda&space;T_G&space;&plus;&space;(\lambda&space;T_G)^2&space;(\rho&space;-&space;\rho^2)&space;\newline&space;R_0&space;=&space;1&space;&plus;&space;\lambda&space;T_G&space;&plus;&space;\rho(1&space;-&space;\rho)&space;(\lambda&space;T_G)^2" target="_blank"><img src="https://latex.codecogs.com/gif.latex?R_0&space;=&space;(1&plus;\frac{\lambda}{\alpha})(1&plus;\frac{\lambda}{\gamma})&space;\newline&space;R_0&space;=&space;(1&plus;\lambda&space;T_L)(1&plus;\lambda&space;T_I)&space;\newline&space;R_0&space;=&space;1&space;&plus;&space;\lambda&space;T_I&space;&plus;&space;\lambda&space;T_L&space;&plus;&space;\lambda^2&space;T_I&space;T_L&space;\newline&space;R_0&space;=&space;1&space;&plus;&space;\lambda&space;(T_I&space;&plus;&space;T_L)&space;&plus;&space;\lambda^2&space;T_I&space;T_G&space;\frac{T_L}{T_G}&space;\newline&space;R_0&space;=&space;1&space;&plus;&space;\lambda&space;T_G&space;&plus;&space;\lambda^2&space;T_I&space;T_G&space;\rho&space;\newline&space;R_0&space;=&space;1&space;&plus;&space;\lambda&space;T_G&space;&plus;&space;\lambda^2&space;\rho(T_G-T_L)T_G&space;\newline&space;R_0&space;=&space;1&space;&plus;&space;\lambda&space;T_G&space;&plus;&space;\lambda^2&space;\rho&space;T_G^2&space;-&space;\lambda^2\rho&space;T_G&space;T_L&space;\newline&space;R_0&space;=&space;1&space;&plus;&space;\lambda&space;T_G&space;&plus;&space;\lambda^2&space;\rho&space;T_G^2&space;-&space;\lambda^2\rho&space;T_G&space;\frac{T_L}{T_G}T_G&space;\newline&space;R_0&space;=&space;1&space;&plus;&space;\lambda&space;T_G&space;&plus;&space;\lambda^2&space;\rho&space;T_G^2&space;-&space;\lambda^2\rho&space;T_G&space;\rho&space;T_G&space;\newline&space;R_0&space;=&space;1&space;&plus;&space;\lambda&space;T_G&space;&plus;&space;\lambda^2&space;\rho&space;T_G^2&space;-&space;\lambda^2&space;\rho^2&space;T_G^2&space;\newline&space;R_0&space;=&space;1&space;&plus;&space;\lambda&space;T_G&space;&plus;&space;(\lambda&space;T_G)^2&space;(\rho&space;-&space;\rho^2)&space;\newline&space;R_0&space;=&space;1&space;&plus;&space;\lambda&space;T_G&space;&plus;&space;\rho(1&space;-&space;\rho)&space;(\lambda&space;T_G)^2" title="R_0 = (1+\frac{\lambda}{\alpha})(1+\frac{\lambda}{\gamma}) \newline R_0 = (1+\lambda T_L)(1+\lambda T_I) \newline R_0 = 1 + \lambda T_I + \lambda T_L + \lambda^2 T_I T_L \newline R_0 = 1 + \lambda (T_I + T_L) + \lambda^2 T_I T_G \frac{T_L}{T_G} \newline R_0 = 1 + \lambda T_G + \lambda^2 T_I T_G \rho \newline R_0 = 1 + \lambda T_G + \lambda^2 \rho(T_G-T_L)T_G \newline R_0 = 1 + \lambda T_G + \lambda^2 \rho T_G^2 - \lambda^2\rho T_G T_L \newline R_0 = 1 + \lambda T_G + \lambda^2 \rho T_G^2 - \lambda^2\rho T_G \frac{T_L}{T_G}T_G \newline R_0 = 1 + \lambda T_G + \lambda^2 \rho T_G^2 - \lambda^2\rho T_G \rho T_G \newline R_0 = 1 + \lambda T_G + \lambda^2 \rho T_G^2 - \lambda^2 \rho^2 T_G^2 \newline R_0 = 1 + \lambda T_G + (\lambda T_G)^2 (\rho - \rho^2) \newline R_0 = 1 + \lambda T_G + \rho(1 - \rho) (\lambda T_G)^2" /></a>

### Parameters
- T<sub>L</sub> is the generation period, which is assumed to be 7.5, taken from reference 1 and 2
- p is the ratio of susceptible (49) turning into confirmed case (59), which is taken to be 0.695, numbers from [People's Daily Weibo](https://m.weibo.cn/u/2803301701)
- Median incubation period is assumed to be 3 days, according to Guan's [Clinical characteristics of 2019 novel coronavirus infection in China](https://www.medrxiv.org/content/10.1101/2020.02.06.20020974v1) (2020)
- rho is the ratio of incubation period over generation period, which is assumed to be 0.4 in the baseline scenario
- R<sub>0</sub> estimation formula refers to reference 3.
- Y<sub>t</sub> `Yt` is the estimated infectious
- Lambda `lamda` is the growth rate of estimated infectious

### Analysis

![R0](https://user-images.githubusercontent.com/56286591/74337661-aafb6480-4ddb-11ea-8a49-9c433ddc81aa.png)

- R<sub>0</sub> exhibits a sharp upward trend from 2020-01-24 to 2020-01-29, rising from 2.24 to 2.51. Since the Wuhan city runs out of medical supply to confirm nCoV cases, the rise of R<sub>0</sub> could be the consequence of outflow of Wuhan residents before lockdown, who seek for medical assistance outside of the Wuhan city and eventually confirmed by the official
- Then, R<sub>0</sub> value fluctuates at 2.50 from 2020-01-30 to 2020-02-03, reaching its peak at 2.51 on 2020-02-03. Time lag between onset of symptons and case confirm is estimated to be 7 to 10 days. The halt of rising trend could be inferred as the consequence of official and social appeal of 'staying at home'. In addition, the fluctuation could be attributed by the inner family infection due to in-house quarantine, where family member infects one and another. 
- Thereafter, R<sub>0</sub> is trending downward, dropping to 2.39 on 2020-02-12. This can be viewed as the positive effect of social and non-pharmeceutical intervention (stay at home and wear facial mask)
- Sensitivity analysis shows that the value of R<sub>0</sub> is highly influenced by the value of T<sub>g</sub>. As more cases being collected, the measure of T<sub>g</sub> will become more accurate

### Sensitivity Analysis
Parameter values
| p | rho | T<sub>g</sub> | Y(t) | R<sub>0</sub> |
| --- | --- | --- | --- | --- |
| 0.695 | 0.4 | 7.5 | 59826 | 2.44 |
| 0.695 | 0.65 | 7.5 | 59826 | 2.42 |
| 0.695 | 0.9 | 7.5 | 59826 | 2.24 |
| 0.695 | 0.4 | 10 | 59826 | 3.05 |
| 0.695 | 0.65 | 10 | 59826 | 3.02 |
| 0.695 | 0.9 | 10 | 59826 | 2.71 |
| 0.8 | 0.4 | 7.5 | 62102 | 2.44 |
| 0.8 | 0.65 | 7.5 | 62102 | 2.42 |
| 0.8 | 0.9 | 7.5 | 62102 | 2.25 |
| 0.8 | 0.4 | 10 | 62102 | 3.06 |
| 0.8 | 0.65 | 10 | 62102 | 3.03 |
| 0.8 | 0.9 | 10 | 62102 | 2.72 |

### Limitation
- This approach assumes the parameter values are the true population parameter values. We have no way to obtain the true value of p, for instance, simply because there is limitation to data access. Even if official data is published, a disproportionate ratio would fail to represent the true parameter value.
- Due to medical shortage, the number of actual confirm cases should be greater than the official report. We would have underestimated the value of Y(t) and have a downward bias in R<sub>0</sub>. 

### Disclaimer
- Data uses API from [BlankerL](https://github.com/BlankerL/DXY-COVID-19-Crawler), which is an infection data realtime crawler. The data source is [Ding Xiang Yuan](https://3g.dxy.cn/newh5/view/pneumonia).

### Reference
1. Li, Q., Guan, X., et al. (2020, January 29). [Early Transmission Dynamics in Wuhan, China, of Novel Coronavirus–Infected Pneumonia](https://www.nejm.org/doi/full/10.1056/NEJMoa2001316#article_references). The New England Journal of Medicine. 

2. Wu, Joseph T., et al. (2020, January 28). [Nowcasting and Forecasting the potential domestic and International Spread of the 2019-nCoV Outbreak Originating in Wuhan, China: a Modeling Study](https://www.thelancet.com/journals/lancet/article/PIIS0140-6736(20)30260-9/fulltext). Lancet.

3. Zhou Tao, Liu, Y., et al. (2020, January 29). [Preliminary Prediction of the Basic Repreduction Number of the Novel Coronavirus 2019-nCoV](http://kns.cnki.net/kcms/detail/51.1656.r.20200204.1640.002.html)

4. Althaus, Christian L. (2014, June). Estimating the Reproduction Number of Ebola Virus (EBOV) During the 2014 Outbreak in West Africa. PLoS Currents. PMC 4169395. PMID 25642364.

5. Gerardo Chowell, Carlos, P., et al. (2004, July). Model Parameters and Outbreak Control for SARS.
