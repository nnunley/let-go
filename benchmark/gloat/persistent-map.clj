(ns bench.persistent-map)

(defn -main []
  (reduce (fn [m i] (assoc m i (* i i)))
          {}
          (range 10000)))
