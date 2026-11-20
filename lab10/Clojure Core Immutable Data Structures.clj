
;; 1. LIST
(def l '(10 20 30))
(println "Original List:" l)
(println "First:" (first l))
(println "Rest:" (rest l))
(println "Conj:" (conj l 5))

;; 2. VECTOR
(def v [10 20 30])
(println "\nOriginal Vector:" v)
(println "Get:" (get v 1))
(println "Count:" (count v))
(println "Conj:" (conj v 40))

;; 3. MAP
(def m {:name "Aman" :age 20})
(println "\nOriginal Map:" m)
(println "Get:" (get m :name))
(println "Assoc:" (assoc m :age 21))
(println "Dissoc:" (dissoc m :age))

;; 4. SET
(def s #{10 20 30})
(println "\nOriginal Set:" s)
(println "Conj:" (conj s 40))
(println "Disj:" (disj s 20))
(println "Contains:" (contains? s 10))
