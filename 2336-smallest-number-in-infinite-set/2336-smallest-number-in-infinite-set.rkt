(require data/heap)

(define smallest-infinite-set%
  (class object%
    (super-new)

    (init-field [hp (make-heap <=)]
                [curr-smallest 0])

    (heap-add! hp 1)

    ; pop-smallest : -> exact-integer?
    (define/public (pop-smallest)
      (let ([ele (heap-min hp)])
        (begin
          (heap-remove-min! hp)
          (when (> ele curr-smallest)
            (set! curr-smallest ele))
          (add-back (add1 curr-smallest))
          ele)))

    ; add-back : exact-integer? -> void?
    (define/public (add-back num)
      (begin
        (heap-remove-eq! hp num)
        (heap-add! hp num)))))