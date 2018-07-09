package main

import(
	"fmt"
)

func main() {
	prime := isPrime(9)
	fmt.Printf("%t\n", prime)
}

func isPrime(num int) bool {
	numberIsPrime := true
	for i := 2; i <= num / 2; i++ {
		if num % i == 0 {
			numberIsPrime = false
			break
		}
	}
	return numberIsPrime
}



// #include <iostream>
// using namespace std;

// int main()
// {
//   int n, i;
//   bool isPrime = true;

//   cout << "Enter a positive integer: ";
//   cin >> n;

//   for(i = 2; i <= n / 2; ++i)
//   {
//       if(n % i == 0)
//       {
//           isPrime = false;
//           break;
//       }
//   }
//   if (isPrime)
//       cout << "This is a prime number";
//   else
//       cout << "This is not a prime number";

//   return 0;
// }