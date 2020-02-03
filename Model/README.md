# 模型

建模，先从简单的传染病模型开始。

# contents

1. basic-fitting and estimation (Imai et al, imperial college london). See [here](https://luxxxlucy.github.io/projects/2020_wuhan/idyll_tryout/index.html)

2. [估计和预测 2019-nCoV 新型冠状病毒在武汉的爆发情况](https://github.com/YiranJing/Coronavirus-Epidemic-2019-nCov)
   - 阶段1: 估计武汉封城时的感染人数(借鉴帝国理工论文)
   - 阶段2: 模拟预测武汉封城后肺炎感染人数以及峰值(借助SIER(susceptible-exposed-infectious- recovered)模型)

***

## 估计和预测 2019-nCoV 新型冠状病毒在武汉的爆发情况
作者: [景怡然](https://github.com/YiranJing)

> 2020年1月23日，交通枢纽的武汉市被封城。900万人民被困在武汉市区。在此之前，有500万人因春节离开武汉。估计机场的国际人流量为1900万。
考虑到新型武汉肺炎的快速传播性和武汉居住人口在封城前后变化巨大，我选择了不同的模型来估计封城前后武汉的感染人数，主要参考和借鉴今日发表的相关论文，数据参考官方数据。

### 模型 1: [估计武汉封城时的感染人数](https://github.com/YiranJing/Coronavirus-Epidemic-2019-nCov/tree/master/Model%201)😷
   - 作者: 景怡然
   - **主要结论： 截止1月23日，武汉有超过 38500 名感染者加确诊者，95%置信区间(30000, 48470)**，根据1月29号海外发现的感染人数计算，引用2018年的交通数据估算。
   > Method: Considering Wuhan is the major air and train transportation hub of China, we use the number of cases exported from Wuhan internationally as the sample, assuming the infected people follow a Possion distribution, then calculate the 95% confidence interval by profile likelihood method. Sensitivity analysis followed by.

   > Reference: [report2 (Jan 21)](https://www.imperial.ac.uk/media/imperial-college/medicine/sph/ide/gida-fellowships/2019-nCoV-outbreak-report-22-01-2020.pdf)

### 模型 2: [模拟预测武汉封城后肺炎感染人数以及峰值](https://github.com/YiranJing/Coronavirus-Epidemic-2019-nCov/blob/master/Model%202/Forecast_Outbreak_Wuhan.ipynb)📈
   - 作者: 景怡然
   > Method: SIER (susceptible-exposed-infectious- recovered) model and Sensitivity analysis

   > Reference: [Nowcasting and forecasting the potential domestic and international spread of the 2019-nCoV outbreak (Jan 31)](https://www.thelancet.com/action/showPdf?pii=S0140-6736%2820%2930260-9)

   - **主要结论:** (根据 2019-12-08 至 2020-02-02 的官方数据)
      - **估计最初的传播速率 R0 (基本传染数) 为: 2.9**
      - **乐观估计武汉肺炎的患者会超过 1.4 万人 (非累计，仅峰值)，峰值最早在2月中旬出现**
      - **考虑到医疗资源不足和官方数据低于实际，武汉肺炎患者的实际峰值可能会在1.6万至2.5万人之间**
      - 封城措施对控制病情有非常显著的作用: 根据模型估算，如果不封城，仅仅隔离患者，武汉患者峰值可能会高达20万。
   - 模型主要假设:
      - 潜伏人群是确诊病例的五倍。(确诊病例按照4109计算，截止2月2日)
      - 23号封城以后，所有确诊病例都会被严格隔离
      - 23号之前，平均1个感染者会传染5个人；23号以后，平均1个感染者最多只会传染1个人
      - 23号之前，武汉人口为1100万；23号后，武汉人口为900万
      - 平均潜伏期为7天，恢复期约为14天
      - 乐观估计医疗资源充足且官方数字准确
![](https://github.com/YiranJing/Coronavirus-Epidemic-2019-nCov/blob/master/Model%202/image/withControl.png)

注释:
- Removed(移除人群): 治愈或者死亡
- Exposed(潜伏人群): 在潜伏期的患者
- Susceptible(易感人群): 健康但有风险被感染的人群
- Infected(确诊并隔离患者): 确诊人群
![](https://github.com/YiranJing/Coronavirus-Epidemic-2019-nCov/blob/master/Model%202/image/SIER2.png)

### 敏感度分析测试
#### 测试1: 用[模型 1](https://github.com/YiranJing/Coronavirus-Epidemic-2019-nCov/tree/master/Model%201)的结论作为初始条件(官方数字低于实际)
  - 在1月23号有38500名感染者，假设其中80%在潜伏期，其余为有明显症状患者
  - 假设目前死亡率等于治愈率，均为3% （根据官方数字）
  - **估计武汉患者最多可超过2.2万人 (非累计，仅峰值)**
#### 测试2: 医疗资源不充足
  - 假设恢复期为20天，而不是14天
  - 估计最初的传播速率 R0 (基本传染数) 为: 3.7
  - **估计武汉肺炎的患者会超过 1.6 万人 (非累计，仅峰值)**, 基于官方数字
#### 测试3: 官方数字低于实际且医疗资源不足
  - 假设恢复期为20天
  - 假设在1月23号有38500名感染者，假设其中80%在潜伏期，其余为有明显症状患者
  - **估计武汉肺炎的患者会超过 2.5 万人 (非累计，仅峰值)**
