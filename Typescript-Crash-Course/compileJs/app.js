"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
let n = 3;
function s(num1, num2) {
    return num1 + num2;
}
console.log(s(n, 3));
function getTotal(numbers) {
    return numbers.reduce((acc, item) => {
        return acc + item;
    }, 0);
}
console.log(getTotal([3, 2, 1]));
const users = {
    firstName: "Md",
    lastName: "Romi",
    role: "Developer",
};
console.log(users.firstName);
const user = {
    name: "MD Romi",
    age: 19,
    role: "Developer",
    address: "123 Dhaka",
};
function login(userData) {
    return userData;
}
console.log(login(user));
const transaction1 = {
    payerAccountNumber: 123,
    payeeAccountNumber: 321,
};
const transaction2 = {
    payerAccountNumber: 123,
    payeeAccountNumber: 456,
};
const bankAccount = {
    accountNumber: 123,
    accountHolder: "Md Romi",
    blance: 4000,
    isActive: true,
    transactions: [transaction1, transaction2],
};
const book = {
    name: "Atomic Habits",
    price: 120,
};
const ebook = {
    name: "Atomic Habits",
    price: 120,
    fileSize: 12,
    format: "pdf",
};
const audioBook = {
    name: "Atomic Habits",
    price: 120,
    fileSize: 12,
    format: "pdf",
    duration: 5,
};
const man = {
    name: "Md Romi",
    age: 19,
};
// Narrowing
function printId(id) {
    if (typeof id === "string")
        console.log(id.toUpperCase());
    console.log(id);
}
printId("romi123");
printId(123);
function getFirstThree(x) {
    return x.slice(0, 3);
}
console.log(getFirstThree("hello"));
console.log(getFirstThree([1, 2, 3, 4, 5]));
// GENERIC
// 1
function logAnyThing(args) {
    console.log(args);
    return args;
}
logAnyThing(["Hello"]);
function getOldest(people) {
    return people.sort((a, b) => b.age - a.age)[0];
}
const people = [{ age: 10 }, { age: 40 }, { age: 30 }];
const players = [
    { name: "Md", age: 30 },
    { name: "Romi", age: 13 },
    { name: "SRD", age: 19 },
];
getOldest(people);
// Assertion
// const person = getOldest(players) as Player
// generics
const person = getOldest(players);
console.log(person.name);
const fetchPostData = (path) => __awaiter(void 0, void 0, void 0, function* () {
    const response = yield fetch(`http://example.com${path}`);
    return response.json();
});
const fetchData = (path) => __awaiter(void 0, void 0, void 0, function* () {
    const response = yield fetch(`http://example.com${path}`);
    return response.json();
});
(() => __awaiter(void 0, void 0, void 0, function* () {
    const posts = yield fetchPostData('/posts');
    const users = yield fetchData('/posts');
    // users[0].age
}))();
function logins(credentials) {
    console.log(credentials);
    return true;
}
const loginUser = {
    username: 'md romi',
    password: '123',
    isAdmin: true
};
logins(loginUser);
const auth = {
    username: 'mdromi',
    password: '123',
    login(username, password) {
        return true;
    }
};
