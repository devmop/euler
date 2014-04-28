import Data.List

main :: IO()
main = putStrLn $ show euler5

euler1 = sum $ filter (multipleOf3 `merge` multipleOf5) [1..999]

euler2 = sum $ filter (even) $ takeWhile (4000000 >) fibs

euler3 = head $ reverse $ primeFactors 600851475143

euler4 = maximum $ filter palindrome threeDigitProducts

euler5 = product $ (simplify.reverse.sort.concat) $ arrangePrimeFactors [2..20]

euler6 = (squareSums 100) - (sumSquares 100)

euler7 = primes !! 10000


arrangePrimeFactors :: [Integer] -> [[(Integer, Int)]]
arrangePrimeFactors = map (countDuplicates.sort.primeFactors)

multipleOf3 :: Integer -> Bool
multipleOf3 x = 3 `divides` x

multipleOf5 :: Integer -> Bool
multipleOf5 x = 5 `divides` x

merge :: (a -> Bool) -> (a -> Bool) -> a -> Bool
merge f g x = (f x) || (g x) 

fibs :: [Integer]
fibs = fibProducer 0 1

fibProducer :: Integer -> Integer -> [Integer]
fibProducer a b = a : (fibProducer b (a+b))

primeFactors n = if prime primes n then [n] else filter (swap divides n) $ takeWhile (\x -> x * x <= n) primes

primes :: [Integer]
primes = 2 : filter (prime primes) [3..]

prime :: [Integer] -> Integer -> Bool
prime [] n = True
prime (x:xs) n
  | x * x > n  = True 
  | x `divides` n = False
  | otherwise = prime xs n

palindrome :: Integer -> Bool
palindrome x = (show x) == (reverse $ show x)

countDuplicates :: [Integer] -> [(Integer, Int)]
countDuplicates [] = []
countDuplicates (x:xs) = let (this, others) = partition (x ==) (x:xs) in (x, length this) : countDuplicates others

threeDigitProducts :: [Integer]
threeDigitProducts = [a*b | a <- [100..999], b <- [100..999]]

simplify :: [(Integer, Int)] -> [Integer]
simplify [] = []
simplify ((x, c) : xs) = (replicate c x) ++ simplify (filter ((x ==).fst) xs)

square :: Integer -> Integer
square x = x * x

swap :: (a -> b -> c) -> b -> a -> c
swap f x y = f y x

sumSquares :: Integer -> Integer
sumSquares n = sum $ map square [1..n]

squareSums :: Integer -> Integer
squareSums n = square $ sum [1..n]

dividesByAll :: [Integer] -> Integer -> Bool
dividesByAll [] n = True
dividesByAll (x:xs) n = if x `divides` n then dividesByAll xs n else False

divides :: Integer -> Integer -> Bool
divides a b = b `mod` a == 0
