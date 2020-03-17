# COVID-19 Epidemic Prevention - Data Science (COVID-19 Open Data)

[简体中文](./README.md) | English

**Goals:**

- Let every newly arrived contributor who has a data-analysis background be able to catch up with the project.
- As a project spawn within the Wuhan2020 community, we are calling for contributors and volunteers who has a data science-related background, including but not limited to Epidemiology, Geographic Information, Economy, etc. We aim at having transparent, open discussions, sharing data and research results that could help people fight against the outbreak caused by the novel Coronavirus.
- Please join [Slack](https://join.slack.com/t/wuhan2020/shared_invite/enQtOTQxMTU4MzgyNTYwLWIxMTMyNWI4NWE2YTk3NGRjZGJhMjUzNmJhMjg1MDQ3OTEzNDE5NGY4MWFhMjRlYWU4MmE3ZGQyOGU4N2YwMzY) to get started. Our channel is **#team-data**

## Important Notice (please read first)

- Contribute by skill set: Please fill out the [SkillSet/TimeZone servey](https://forms.gle/Yyh1DY5EkzqCpA897) so people could coordinate online or offline more easily. Please also see [Info on Google Docs](https://docs.google.com/spreadsheets/d/16r_sJtGOVuU0pNulKmqXrvrm3N5mxLb_b4zgcdi_C1c/edit?usp=sharing)

- **Risk Disclosure**：Predictions are not always right and could cuase serious problems, which is tricky and risky. Currently this porjects mainly focus on:
   - Establishing professional knowledge base
   - Basic database/data repo
   - Data Visualization
   - Rapidly testing models

Please use your own judgement to evaluate the consequences and social impact of your predictions or modeling, if you are planning to do so!

## **List of Resources**

- [John Hopkins Dashboard](http://gisanddata.maps.arcgis.com/apps/opsdashboard/index.html) Geographical visualization of the epidemic
- [Data Sources](https://docs.google.com/spreadsheets/d/1JNQnFYJpR7PxQo5K5lwXTuE0F6jprhMXuy7DPnV9H90/edit?usp=sharing) Including official data sources and 3rd-party data. *Openly Editable*
- [Modeling Navigation](https://docs.google.com/document/d/1Gm1xBevUXdTj5ZnMK82rKW7EblQoB8ML5OhIYMHmm1U/edit?usp=sharing) Share the models and references here, *Openly Editable*
- [Visualization and epidemic papers and references](https://docs.google.com/document/d/17v0cAcyhm2Yd0FNqmgHxcVNR2nZdQ7LmSLfxt21wv5w) Including links to Elsevier, Wiley, Lancet, etc，*Openly Editable*

### Other 3rd-party Resources

- [Outbreak analytics](https://royalsocietypublishing.org/doi/pdf/10.1098/rstb.2018.0276) Application of data science in epidemic outbreak (provided by Eric Y.L.)
- [患者相同行程查询工具](患者相同行程查询工具)

## **Resources provided by Individual Contributors**

Every contributor is welcome to edit this section by opening a PR.

### Data Source：

- Wangxin：[Collection](https://docs.google.com/spreadsheets/d/1M-iFpgweAAyRUriT_NsmPdkXCcDFReE3-8a0fYNQQc4/edit#gid=235095609)
- Stockard：[Summary based on official released text data of the outbreak](https://docs.google.com/spreadsheets/d/1WnfMY1M3k8x9WHPGpx_jZc2EbSFQv7rEwHj5lJdBnVo/edit?usp=sharing)
- Yilun Guan: Data of the population traveled out from Wuhan during Jan.1 to Jan.26 by Baidu
- Michael: [Summary of the news and repots about the novel-Coronavirus in North America, on Telegram](https://t.me/ncov_2019_us)
- Guangchuang Yu: [nCov2019: R package to query for recent and historical data](https://github.com/GuangchuangYu/nCov2019)

### 模型和调试：

- Tongshuang Wu(Vegalite)：[Online tool for data wrangling](https://idyll-lang.org/)
- Jialin Lu ：[Reproduced model of the estimation model from Imperial College London](https://luxxxlucy.github.io/projects/2020_wuhan/index.html)
- Yiran Jing ：[Estimation and prediction of the outbreak in Wuhan](https://github.com/YiranJing/Coronavirus-Epidemic-2019-nCov)
   - Phase1: Estimation of the infected population in Wuhan (refer to the ICL model)
   - Phase2: Simulate and predict the infected population and peak after the shutdown of Wuhan. (Leverage the SIER(susceptible-exposed-infectious- recovered) method)

## **Contributing Guide**

**Ways of contributing**
- Contribute your data to `Data/` folder.
- Sharing your model to `/Model/`.
- Submit the references and related papers to `Reference/`.
- Get involved in the development work, including data wrangling, API development, modeling, integrate with the main project, etc.

Please download [Github Desktop](https://desktop.github.com/) and get yourself familiar with the guide and tutorials before you clone this repository by either using the "Clone or Download" button on this page, if you haven't used Github before.

Feel free to join [Trello](https://trello.com/invite/b/fwW7gZUY/87d9ee972226d4d13e59ed3058c69266/ncovdata-team) to view and pick the tickets from the backlog.

Author: @Stockard, co-edit @TensorFrozen
