var Cat = new Object()

Cat.name = "五仁"
Cat.age = 1
Cat.sex = 'f'
console.log(Cat);



var Dog = {
    name : "大黄",
    age : 2,
    sex : 'm'
}
console.log(Dog);


function creatMouse(name,age,sex){
    this.name = name;
    this.age = age;
    this.sex = sex;
}
Mouse = new creatMouse("米瑞",1,'f')
console.log(Mouse)


function creatTiger(name,age,sex) {
    tiger = {};
    tiger.name = name;
    tiger.age = age;
    tiger.sex = sex;
    return tiger;
}
Tiger = creatTiger("大虎",2,'m')
console.log(Tiger)