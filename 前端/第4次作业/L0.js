const stu = {
    name: 'MING',
    hobby: ['play', 'run', 'sing'],
    address: {
        school: 'ChongQing',
        home: 'HENAN'
    },
    title: ['student',{year:2022}],
    skills: {
        speak() {
            this.name = 'JACK';
        }
    }
}
let {name,hobby:[h1,h2,h3],address:{school,home:HOME},title:[,{year}]} = stu
console.log(name)
console.log(h1,h2,h3)
console.log(HOME)
console.log(year)