# SEIR Forecast

#### Author: Yijun Wang, Owen Xu
#### Date: 2020-02-14
#### Last Update: 2020-02-16

### Usage
- Download my Jupyter notebook file: [SEIR.ipynb](https://github.com/yijunwang0805/YijunWang/blob/master/SEIR%20Forecast_Yijun%20Wang%20%26%20Owen%20Xu/SEIR.ipynb).
- First part of the code will load data from an API connection, which is provided by [BlankerL](https://lab.isaaclin.cn/nCoV/). 
- ```R0Func()``` is the function that calculates the newest basic reproduction number given up to date statistics. Its ```inputs``` are the number of ```confirm``` cases, the number of ```suspect``` cases, and days ```t``` since the start of the epidemic, December 1st, 2019.
- ```SEIR()``` is the epidemic model that describes the system of differential equations.
- ```betaFunc()``` and ```gammaFunc()``` calculate the value of transmissibility and removal rate, respectively.
- ```spi.odeint()``` solves the system of differential equations. Its ```inputs``` are the epidemic model ```SEIR()```, initial value of susceptible, exposed, infectious, removal ```INI```, and the number of days since the epidemic ```Time```

### Summary
- This study seeks to forecast the peak time of SARS-CoV-2 cases. We find, under the assumptions of no quaratine intervention, Wuhan reach peak infectiouson March 3, 2020; Beijing, Shanghai, and Guangzhou would each peak infectious in the middle of May. Sensitivity analysis shows that reducing half of the number of catchment size and the reproductive number would reduce the magnitude of epidemic by more than 60%, while lengthening the peak to June and duration of the epidemic to August.

### Background
- There has been a novel coronavirus (2019-nCoV) pneumonia outbreak in Wuhan, China since December 2019 which spreads domestically and internationally. 
- On 2019-12-01, the first case of nCoV exhibits symptons, according to Huang's [Clinical features of patients infected with 2019 novel coronavirus in Wuhan, China](https://www.thelancet.com/journals/lancet/article/PIIS0140-6736(20)30183-5/fulltext#seccestitle10). This date is used as the first day of the epidemic for Wuhan city.
- On 2020-1-23, Wuhan city lockdowns [(ifeng news)](http://news.ifeng.com/c/7tpL47zV2Vy). Before Wuhan lockdown, 5 million people left the city, leaving 9 million in town [(Tencent News)](https://new.qq.com/sv1/qd/aoyou.html?cmsid=20200127A0EFXJ00). 
- There is a lag in time between onset of symptons and case confirmed, particularly due to staff and medical supply shortage. The median time from onset of symptoms to first hospital admission was 7 days, according to Huang's [Clinical featupracticeres of patients infected with 2019 novel coronavirus in Wuhan, China](https://www.thelancet.com/journals/lancet/article/PIIS0140-6736(20)30183-5/fulltext#seccestitle10) (2020).
- After the Wuhan outbreak, people quaratine at home if their cases are not severe, which is a common  to avoid cross-infection. 

### Assumption
- Assume homogeneous features across cities
- Assume that infected individuals were not infectious during the incubation period (Wu 2020; Zhou, 2020)
- Assume each city as a closed environment
- Assume population growth rate and death rate are zero
- Assume people exhibit consistent behaviors before and during the epidemic
- Assume the number of initial infectious takes the value of the confirm counts
- Assume no quarantine or other mitigation intervention is implemented

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

### Parameters
- The value of```R0``` takes the newest estimation from [Estimation of R0.ipynb](https://github.com/yijunwang0805/YijunWang/blob/master/Estimation%20of%20R0_Yijun/Estimation%20of%20R0.ipynb) as of writing this study Feb 15, 2020, while research varies from 1.4 to 3.9 (WHO, 2020; Zhou, 2020; Read, 2020).
  * [Zhao](https://www.ijidonline.com/article/S1201-9712(20)30053-9/fulltext) estimates that the mean R0 ranges from 2.24 to 3.58.
  * [WHO](https://www.who.int/news-room/detail/23-01-2020-statement-on-the-meeting-of-the-international-health-regulations-(2005)-emergency-committee-regarding-the-outbreak-of-novel-coronavirus-(2019-ncov))'s preliminary estimate of R<sub>0</sub> is in the range of 1.4 to 2.5. 
  * [Zhou, et al](https://arxiv.org/abs/2001.10530) estimates basic reproduction number falls between 2.8 to 3.9. 
  * Read, et al's [Novel coronavirus 2019-nCoV: early estimation of epidemiological parameters and epidemic predictions](https://www.medrxiv.org/content/10.1101/2020.01.23.20018549v2.article-info) estimates R<sub>0</sub> to be 3.11.
- City population parameter ```N``` refers to [The Economic Observer](https://baijiahao.baidu.com/s?id=1656943894281117716&wfr=spider&for=pc).
- The date Wuhan implemented lockdown, January 23, 2020, is used as the first day of the forecast for all the city except Wuhan.
- The first case is documented on December 1, 2019 in Wuhan, which is used as the first day of the forecast for Wuhan. 
- The number of initial infectious takes the value of confirm counts on Jan 23, 2020.

### Analysis

![wuhan](https://user-images.githubusercontent.com/56286591/74587420-15571380-502d-11ea-955a-c6869b693af4.png)

**Forecast Infectious**
| Cities | Forecast Infectious Counts at Peak Time | Peak Date | 
| --- | --- | --- | 
| Beijing | 3,943,645 | 2020-05-11 | 
| Shanghai | 5,213,689 | 2020-05-16 | 
| Guangzhou | 4,241,083  | 2020-05-21 | 
| Wuhan | 1,734,367 | 2020-03-03 |

Assuming similar transmissibility in all the cities, Shanghai will have the largest magnitude of epidemic while Wuhan will have the earliest peak time.

**Sensitivity Analysis**
| Wuhan | N/2 | I(0)/2 | beta/2
| --- | --- | --- | --- |
| Forecast Infectious Counts at Peak Time| 335,692 | 1,734,473 | 660,246
| Peak Date | 2020-06-10 | 2020-03-09 | 2020-06-18

If transmissibility ```beta``` was reduced by 50% in all cities, both the growth rate and magnitude of Wuhan epidemic would be substantially reduced. The epidemic peak would be delayed by about 2 months and its magnitude reduced by about 60%. The reduction in
transmissibility would push the viral reproductive number to about 1.2, in which case the epidemic would grow slowly without peaking during the first half of 2020.

Similarly, a 50% reduction in the catchment size ```N``` would reduce the magnitude of the epidemic by 80%, controlling for all the other variables. The epidemic peak would also be prolonged by approximately 2 months. 

Unlike transmissibility and catchment size, our estimates suggested that a 50% reduction in initial infectious ```I(0)``` would have a negligible effect on epidemic dynamics.

### Limitation
- Our findings are critically dependent on the assumptions underpinning our model, and the timing and reporting of confirm cases. It is also vulnerable to variation in medical supply, heterogeneity across cities, and consistent behavior during the epidemic.

### Disclaimer
- Data uses API from [BlankerL](https://github.com/BlankerL/DXY-COVID-19-Crawler), which is an infection data realtime crawler. The data source is [Ding Xiang Yuan](https://3g.dxy.cn/newh5/view/pneumonia).

### Reference
1. Li, Q., Guan, X., et al. (2020, January 29). [Early Transmission Dynamics in Wuhan, China, of Novel Coronavirus–Infected Pneumonia](https://www.nejm.org/doi/full/10.1056/NEJMoa2001316#article_references). The New England Journal of Medicine. 

1. Wu, Joseph T., et al. (2020, January 28). [Nowcasting and Forecasting the potential domestic and International Spread of the 2019-nCoV Outbreak Originating in Wuhan, China: a Modeling Study](https://www.thelancet.com/journals/lancet/article/PIIS0140-6736(20)30260-9/fulltext). Lancet.

1. Zhou Tao, Liu, Y., et al. (2020, January 29). [Preliminary Prediction of the Basic Repreduction Number of the Novel Coronavirus 2019-nCoV](http://kns.cnki.net/kcms/detail/51.1656.r.20200204.1640.002.html)

1. Althaus, Christian L. (2014, June). Estimating the Reproduction Number of Ebola Virus (EBOV) During the 2014 Outbreak in West Africa. PLoS Currents. PMC 4169395. PMID 25642364.

1. Gerardo Chowell, Carlos, P., et al. (2004, July). Model Parameters and Outbreak Control for SARS.

1. [Population](http://worldpopulationreview.com/world-cities/). (2019-05-12). Retrieved 2020-02-14

1. Read JM, Bridgen JRE, Cummings DAT, Ho A, Jewell CP. [Novel coronavirus
2019-nCoV: early estimation of epidemiological parameters and epidemic predictions](https://www.medrxiv.org/content/10.1101/2020.01.23.20018549v1.full.pdf).
medRxiv. 2020:2020.01.23.20018549. doi: 10.1101/2020.01.23.20018549.

1. [Doraemon](https://zhuanlan.zhihu.com/p/104091330?utm_source=wechat_session&utm_medium=social&utm_oi=1166139227388686336)'s code

1. [Udacity](https://classroom.udacity.com/courses/cs222)'s code


