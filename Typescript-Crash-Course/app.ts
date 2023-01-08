let n = 3;

function s(num1: number, num2: number) {
  return num1 + num2;
}

console.log(s(n, 3));

function getTotal(numbers: Array<number>) {
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

// TYPE ALIAS
type User = {
  name: string;
  age: number;
  role: string;
  address?: string;
};

const user: User = {
  name: "MD Romi",
  age: 19,
  role: "Developer",
  address: "123 Dhaka",
};

function login(userData: User): User {
  return userData;
}
console.log(login(user));

// Interfaces
interface Transaction {
  payerAccountNumber: number;
  payeeAccountNumber: number;
}

interface BankAccount {
  accountNumber: number;
  accountHolder: string;
  blance: number;
  isActive: boolean;
  transactions: Transaction[];
}

const transaction1: Transaction = {
  payerAccountNumber: 123,
  payeeAccountNumber: 321,
};

const transaction2: Transaction = {
  payerAccountNumber: 123,
  payeeAccountNumber: 456,
};

const bankAccount: BankAccount = {
  accountNumber: 123,
  accountHolder: "Md Romi",
  blance: 4000,
  isActive: true,
  transactions: [transaction1, transaction2],
};

// Extend
interface Book {
  name: string;
  price: number;
}

interface EBook extends Book {
  fileSize: number;
  format: string;
}

interface AudioBook extends EBook {
  duration: number;
}

const book: Book = {
  name: "Atomic Habits",
  price: 120,
};

const ebook: EBook = {
  name: "Atomic Habits",
  price: 120,
  fileSize: 12,
  format: "pdf",
};

const audioBook: AudioBook = {
  name: "Atomic Habits",
  price: 120,
  fileSize: 12,
  format: "pdf",
  duration: 5,
};

// Marge
interface Man {
  name: string;
}
interface Man {
  age: number;
}

const man: Man = {
  name: "Md Romi",
  age: 19,
};

// Unions
type ID = number | string;

// Narrowing
function printId(id: ID) {
  if (typeof id === "string") console.log(id.toUpperCase());
  console.log(id);
}
printId("romi123");
printId(123);

function getFirstThree(x: string | number[]) {
  return x.slice(0, 3);
}
console.log(getFirstThree("hello"));
console.log(getFirstThree([1, 2, 3, 4, 5]));

// GENERIC

// 1
function logAnyThing<T>(args: T): T {
  console.log(args);
  return args;
}
logAnyThing(["Hello"]);

// 2
interface HasAge {
  age: number;
}
interface Player {
  name: string;
  age: number;
}

function getOldest<T extends HasAge>(people: T[]): T {
  return people.sort((a, b) => b.age - a.age)[0];
}

const people: HasAge[] = [{ age: 10 }, { age: 40 }, { age: 30 }];
const players: Player[] = [
  { name: "Md", age: 30 },
  { name: "Romi", age: 13 },
  { name: "SRD", age: 19 },
];

getOldest(people);

// Assertion
// const person = getOldest(players) as Player

// generics
const person = getOldest(players)
console.log(person.name);


interface IPost {
    title: string;
    id: number;
    description: string;
}

interface IUser {
    id: number;
    name: string;
    age: number;
}

const fetchPostData = async (path: string): Promise<IPost[]> => {
    const response = await fetch(`http://example.com${path}`) ;
    return response.json()
}

const fetchData = async <ResultType>(path: string): Promise<ResultType[]> => {
    const response = await fetch(`http://example.com${path}`) ;
    return response.json()
}

(async () => {
    const posts = await fetchPostData('/posts')
    const users = await fetchData<IUser>('/posts')
    // users[0].age
})()


// Structural Typing / Duck Typing
interface ICredentials {
    username: string;
    password: string;
    isAdmin: boolean
}

function logins(credentials: ICredentials):boolean {
    console.log(credentials);
    return true
}

const loginUser = {
    username: 'md romi',
    password: '123',
    isAdmin: true
}
logins(loginUser)

interface IAuth {
    username: string;
    password: string;
    login(username: string, password: string): boolean
}

const auth: IAuth = {
    username: 'mdromi',
    password: '123',
    login(username: string, password: string) {
        return true
    }
}