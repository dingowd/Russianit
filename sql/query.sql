select Сотрудник, max(q0) as "1.8", max(q1) as "2.8", max(q2) as "3.8", max(q3) as "4.8", max(q4) as "5.8", max(q5) as "6.8", max(q6) as "7.8", max(q7) as "8.8", max(q8) as "9.8", max(q9) as "10.8", max(q10) as "11.8", max(q11) as "12.8", max(q12) as "13.8", max(q13) as "14.8", max(q14) as "15.8", max(q15) as "16.8", max(q16) as "17.8", max(q17) as "18.8", max(q18) as "19.8", max(q19) as "20.8", max(q20) as "21.8", max(q21) as "22.8", max(q22) as "23.8", max(q23) as "24.8", max(q24) as "25.8", max(q25) as "26.8", max(q26) as "27.8", max(q27) as "28.8", max(q28) as "29.8", max(q29) as "30.8", max(q30) as "31.8"
from
(
select EmployeeID as Сотрудник,
if(DATE(StartPeriod) = '2021-08-01' or DATE(EndPeriod) = '2021-08-01', '+', '') as q0,
if(DATE(StartPeriod) = '2021-08-02' or DATE(EndPeriod) = '2021-08-02', '+', '') as q1,
if(DATE(StartPeriod) = '2021-08-03' or DATE(EndPeriod) = '2021-08-03', '+', '') as q2,
if(DATE(StartPeriod) = '2021-08-04' or DATE(EndPeriod) = '2021-08-04', '+', '') as q3,
if(DATE(StartPeriod) = '2021-08-05' or DATE(EndPeriod) = '2021-08-05', '+', '') as q4,
if(DATE(StartPeriod) = '2021-08-06' or DATE(EndPeriod) = '2021-08-06', '+', '') as q5,
if(DATE(StartPeriod) = '2021-08-07' or DATE(EndPeriod) = '2021-08-07', '+', '') as q6,
if(DATE(StartPeriod) = '2021-08-08' or DATE(EndPeriod) = '2021-08-08', '+', '') as q7,
if(DATE(StartPeriod) = '2021-08-09' or DATE(EndPeriod) = '2021-08-09', '+', '') as q8,
if(DATE(StartPeriod) = '2021-08-10' or DATE(EndPeriod) = '2021-08-10', '+', '') as q9,
if(DATE(StartPeriod) = '2021-08-11' or DATE(EndPeriod) = '2021-08-11', '+', '') as q10,
if(DATE(StartPeriod) = '2021-08-12' or DATE(EndPeriod) = '2021-08-12', '+', '') as q11,
if(DATE(StartPeriod) = '2021-08-13' or DATE(EndPeriod) = '2021-08-13', '+', '') as q12,
if(DATE(StartPeriod) = '2021-08-14' or DATE(EndPeriod) = '2021-08-14', '+', '') as q13,
if(DATE(StartPeriod) = '2021-08-15' or DATE(EndPeriod) = '2021-08-15', '+', '') as q14,
if(DATE(StartPeriod) = '2021-08-16' or DATE(EndPeriod) = '2021-08-16', '+', '') as q15,
if(DATE(StartPeriod) = '2021-08-17' or DATE(EndPeriod) = '2021-08-17', '+', '') as q16,
if(DATE(StartPeriod) = '2021-08-18' or DATE(EndPeriod) = '2021-08-18', '+', '') as q17,
if(DATE(StartPeriod) = '2021-08-19' or DATE(EndPeriod) = '2021-08-19', '+', '') as q18,
if(DATE(StartPeriod) = '2021-08-20' or DATE(EndPeriod) = '2021-08-20', '+', '') as q19,
if(DATE(StartPeriod) = '2021-08-21' or DATE(EndPeriod) = '2021-08-21', '+', '') as q20,
if(DATE(StartPeriod) = '2021-08-22' or DATE(EndPeriod) = '2021-08-22', '+', '') as q21,
if(DATE(StartPeriod) = '2021-08-23' or DATE(EndPeriod) = '2021-08-23', '+', '') as q22,
if(DATE(StartPeriod) = '2021-08-24' or DATE(EndPeriod) = '2021-08-24', '+', '') as q23,
if(DATE(StartPeriod) = '2021-08-25' or DATE(EndPeriod) = '2021-08-25', '+', '') as q24,
if(DATE(StartPeriod) = '2021-08-26' or DATE(EndPeriod) = '2021-08-26', '+', '') as q25,
if(DATE(StartPeriod) = '2021-08-27' or DATE(EndPeriod) = '2021-08-27', '+', '') as q26,
if(DATE(StartPeriod) = '2021-08-28' or DATE(EndPeriod) = '2021-08-28', '+', '') as q27,
if(DATE(StartPeriod) = '2021-08-29' or DATE(EndPeriod) = '2021-08-29', '+', '') as q28,
if(DATE(StartPeriod) = '2021-08-30' or DATE(EndPeriod) = '2021-08-30', '+', '') as q29,
if(DATE(StartPeriod) = '2021-08-31' or DATE(EndPeriod) = '2021-08-31', '+', '') as q30
from timework
) as query1
group by Сотрудник;