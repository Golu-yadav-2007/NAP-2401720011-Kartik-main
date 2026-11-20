
;; First Collection
(def nums [1 2 3 4 5])

;; 1. MAP - Square numbers
(println "Map:" (map #(* % %) nums))

;; 2. FILTER - Even numbers
(println "Filter:" (filter even? nums))

;; 3. REDUCE - Sum
(println "Reduce:" (reduce + nums))

;; 4. TAKE and DROP
(println "Take:" (take 3 nums))
(println "Drop:" (drop 2 nums))

;; Threading Macro
(println "Threading:"
  (->> nums
       (filter even?)
       (map #(* % %))))

;; Second Collection
(def nums2 [10 20 30 40])

(println "\nSecond Collection:")
(println "Map:" (map #(* % %) nums2))
(println "Filter:" (filter even? nums2))
(println "Reduce:" (reduce + nums2))
(println "Take:" (take 2 nums2))
(println "Drop:" (drop 1 nums2))
