(ns bench.loop-recur)

(defn -main []
  (loop [i 0 acc 0]
    (if (< i 1000000)
      (recur (inc i) (+ acc i))
      acc)))
