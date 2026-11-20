(ns exp23.short)

;; (2) Custom exception with ex-info
(defn check-amt [amt]
  (when (neg? amt)
    (throw (ex-info "Negative amount" {:amt amt})))
  amt)

;; (1) & (3) try, catch, finally + ex-message & ex-data recovery
(defn demo [x]
  (try
    (case x
      :arith (/ 10 0)
      :npe   (.length nil)
      :ex    (check-amt -50)
      :ok    (check-amt 100))
    (catch ArithmeticException _ 0)
    (catch NullPointerException _ "")
    (catch clojure.lang.ExceptionInfo e
      (println "Err:" (ex-message e) "| Data:" (ex-data e))
      0)
    (finally (println "Done."))))

(defn -main []
  (doseq [m [:arith :npe :ex :ok]]
    (println m "->" (demo m))))
