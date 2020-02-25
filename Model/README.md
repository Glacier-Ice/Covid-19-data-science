# 模型

建模，先从简单的传染病模型开始。

# contents

1. basic-fitting and estimation (Imai et al, imperial college london)
   - See [here](https://luxxxlucy.github.io/projects/2020_wuhan/idyll_tryout/index.html) for visualization

2. [估计和预测 2019-nCoV 新型冠状病毒在武汉的爆发情况](https://github.com/YiranJing/Coronavirus-Epidemic-2019-nCov)
   - 阶段1: 估计武汉封城时的感染人数(借鉴帝国理工论文)
   - 阶段2: 模拟预测武汉封城后肺炎感染人数以及峰值(借助SIER(susceptible-exposed-infectious- recovered)模型)

3. Estimating peak infectious time in metropolis cities 预估COVID在各大城市的爆发情况
   - 阶段1: estimate the basic reproduction number by deriving R<sub>0</sub> from the SEIR model. As of 2020-02-14, R<sub>0</sub> is estimated to be 2.41.
   - 阶段2: estimate the peak time of infectious using SEIR model.
