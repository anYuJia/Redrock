import random
hwin = ywin = 0
dict=["石头","剪刀","布"]
while ywin<3 and hwin<3:
    random_num=random.randint(1,100)
    print("当前比分 你 : yyz =%d : %d"%(ywin,hwin))
    yourhand=input("请输入你的手法(剪刀，石头，布)：")
    if yourhand not in dict:
        print("请输入正确手法，小心我来锤你QAQ")
        continue
    if random_num%3==0:
        hishand="石头"
    elif random_num%3==2:
        hishand="布"
    else:
        hishand="石头"
    if yourhand=="布":
        print("你出了布")
        if hishand=="剪刀":
            print("yyz出了剪刀")
            hwin+=1
        elif hishand=="石头":
            print("yyz出了石头")
            ywin+=1
        elif hishand=="布":
            print("yyz出了布")
            continue
    if yourhand=="石头":
        print("你出了石头")
        if hishand=="布":
            print("yyz出了布")
            hwin+=1
        elif hishand=="剪刀":
            print("yyz出了剪刀")
            ywin+=1
        elif hishand=="石头":
            print("yyz出了石头")
            continue
    if yourhand=="剪刀":
        print("你出了剪刀")
        if hishand=="石头":
            print("yyz出了石头")
            hwin+=1
        elif hishand=="布":
            print("yyz出了布")
            ywin+=1
        elif hishand=="剪刀":
            print("yyz出了剪刀")
            continue
print("当前比分 你 : yyz =%d : %d"%(ywin,hwin))
if ywin>hwin:
    print("你战胜了yyz，你好棒")
else:
    print("你不行啊，你居然输了，加油，奥里给！")