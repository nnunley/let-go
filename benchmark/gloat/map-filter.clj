(ns bench.map-filter)

(defn -main []
  (reduce + 0
    (take 100
      (filter even?
        (map #(* % %) (range 10000))))))
