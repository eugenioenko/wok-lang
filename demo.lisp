(print (+ 1 2 3))
(print (- 2))
(print (- 8 4))
(print (print "yes"))
(cond
  (0 (print "0"))
  (0 (print "1"))
  (2 (print "2"))
)

(if 1
  (print "yes")
  (print "no")
)

(print "hello world")
(:= index 0)
(print index)
(while (! (== index 10))
  (print
    (:= index (+ index 1))
  )
)


(defun function (alpha beta)
  (print alpha)
  (return-from "function" "the_return_value")
  (print beta)
)

(function "hello" "world")
(print (function "first" "second"))

