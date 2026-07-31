// function add(x) {
//     return function(y) {
//         return x + y;
//     };
// }
const add = x => y => x + y;

console.log(add(1)(2)); //Output: 3
let increment = add(1);
console.log(increment(2)); //Output: 3
