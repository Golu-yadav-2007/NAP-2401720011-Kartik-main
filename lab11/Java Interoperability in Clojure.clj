
;; 1. Java Static Methods
(println "Absolute:" (Math/abs -10))
(println "Current Time:" (System/currentTimeMillis))

;; 2. Create Java Objects
(def arr (java.util.ArrayList.))
(def mp (java.util.HashMap.))

;; 3. Instance Methods
(.add arr "Aman")
(.add arr "Rahul")
(println "ArrayList:" arr)

(.put mp "Name" "Aman")
(.put mp "Age" 20)
(println "HashMap:" mp)

;; 4. Convert Clojure to Java
(def nums [10 20 30])
(def javaList (java.util.ArrayList. nums))
(println "Java List:" javaList)

;; Convert Java to Clojure
(println "Clojure List:" (vec arr))
