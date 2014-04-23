main :: IO()
main = putStrLn $ show euler6


euler1 = sum $ filter (multipleOf3 `merge` multipleOf5) [1..999]

euler2 = sum $ filter (even) $ takeWhile (4000000 >) fibs

euler3 = head $ reverse $ primeFactors 600851475143

euler4 = maximum $ filter palindrome threeDigitProducts

euler5 = (squareSums 100) - (sumSquares 100)

euler6 = primes !! 10001

multipleOf3 :: Integer -> Bool
multipleOf3 x = x `mod` 3 == 0

multipleOf5 :: Integer -> Bool
multipleOf5 x = x `mod` 5 == 0

merge :: (a -> Bool) -> (a -> Bool) -> a -> Bool
merge f g x = (f x) || (g x) 

fibs :: [Integer]
fibs = fibProducer 0 1

fibProducer :: Integer -> Integer -> [Integer]
fibProducer a b = a : (fibProducer b (a+b))

primeFactors n = filter (\x -> n `mod` x == 0) $ takeWhile (\x -> x * x <= n) primes

primes :: [Integer]
primes = 2 : filter (prime primes) [3..]

prime :: [Integer] -> Integer -> Bool
prime [] n = True
prime (x:xs) n
  | x * x > n  = True 
  | n `mod` x == 0 = False
  | otherwise = prime xs n

palindrome :: Integer -> Bool
palindrome x = (show x) == (reverse $ show x)

threeDigitProducts :: [Integer]
threeDigitProducts = [a*b | a <- [100..999], b <- [100..999]]

square :: Integer -> Integer
square x = x * x

sumSquares :: Integer -> Integer
sumSquares n = sum $ map square [1..n]

squareSums :: Integer -> Integer
squareSums n = square $ sum [1..n]