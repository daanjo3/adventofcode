const fs = require('fs')

let monkeys = readMonkeys('monkey-short.json')
const supermod = monkeys
    .filter((m) => m.op.n != 'old')
    .map((m) => parseInt(m.op.n))
    .reduce((tot, cur) => tot * cur, 1)

function readMonkeys(filename) {
    const rawdata = fs.readFileSync(filename);
    return JSON.parse(rawdata);
}

function inspectItem(monkey, item) {
    if (monkey.inspectCount == undefined) {
        monkey.inspectCount = 0
    }
    monkey.inspectCount++

    const { op, n } = monkey.op
    let value = 0
    if (n == 'old') {
        value = item
    } else {
        value = parseInt(n)
    }
    if (op == '*') {
        return item * value
    }
    if (op == '+') {
        return item += value
    }
}

function relief(item) {
    return Math.floor(item / 3.0)
}

function throwItem(monkeys, monkey, item) {
    const { test, pos, neg } = monkey.test
    const throwTo = item % test == 0 ? pos : neg
    monkeys[throwTo].items.push(item)
}

function doMonkeyAction(monkeys, monkey) {
    while (monkey.items.length > 0) {
        let item = monkey.items.shift() % supermod
        item = inspectItem(monkey, item)
        // item = relief(item)
        throwItem(monkeys, monkey, item)
    }
}

async function part1() {
    console.log('start 11p1')
    // console.log(JSON.stringify(monkeys, null, 2))
    for (i = 0; i < 10000; i++) {
        monkeys.forEach((m) => doMonkeyAction(monkeys, m))
        if ((i + 1) <= 10 || (i + 1) % 1000 == 0 ){
            console.log(`After round ${i+1}, the monkeys are holding items with these worry levels:`)
            monkeys.map((m) => `Monkey ${m.id}: ${m.items.join(',')}`).forEach((s) => console.log(s))
        }
    }
    monkeys.forEach((m) => console.log(`monkey ${m.id} inspected items ${m.inspectCount} times.`))

    const assholeMonkeys = monkeys.sort((a, b) => b.inspectCount - a.inspectCount).slice(0, 2)
    console.log(`Level of monkey business is ${assholeMonkeys[0].inspectCount*assholeMonkeys[1].inspectCount}.`)

    console.log('end 11p1')
}

part1()
