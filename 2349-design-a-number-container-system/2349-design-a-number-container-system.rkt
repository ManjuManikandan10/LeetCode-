(require racket/class)
(require data/heap)
(define number-containers%
  (class object%
    (super-new)
    
    (init-field [hash:number.heap
                 (make-hash)]
                [hash:index.number
                 (make-hash)])
    (define number.heap
      hash:number.heap)
    (define index.number
      hash:index.number)
    (define/public (change index
                           number)
      (let*
          ((position_value:number
            (hash-ref index.number
                      index
                      #f))
           (previous:heap:index
            (if position_value:number
                (hash-ref number.heap
                          position_value:number)
                #f))
           (current:heap:index
            (hash-ref number.heap
                      number
                      #f)))
        (cond
          ((not previous:heap:index)
           (if current:heap:index
               (begin
                 (heap-add! current:heap:index
                            index)
                 (hash-set! index.number
                            index
                            number))
               (let
                   ((new_heap:heap:index
                     (make-heap <)))
                 (begin
                   (heap-add! new_heap:heap:index
                              index)
                   (hash-set! index.number
                              index
                              number)
                   (hash-set! number.heap
                              number
                              new_heap:heap:index)))))
          ((not current:heap:index)
           (let
               ((new_heap:heap:index
                 (make-heap <)))
             (begin
               (heap-remove-eq! previous:heap:index
                                index)
               (heap-add! new_heap:heap:index
                          index)
               (hash-set! number.heap
                          number
                          new_heap:heap:index)
               (hash-set! index.number
                          index
                          number))))
          ((equal? current:heap:index
                   previous:heap:index)
           values)
          (else
           (begin
             (heap-remove-eq! previous:heap:index
                              index)
             (heap-add! current:heap:index
                        index)
             (hash-set! index.number
                        index
                        number))))))
 
    (define/public (find number)
      (let*
          ((relevant:heap:index
            (hash-ref number.heap
                      number
                      #f))
           (heap_check:integer
            (if (and relevant:heap:index
                     (positive? (heap-count relevant:heap:index)))
                (heap-min relevant:heap:index)
                -1)))
        heap_check:integer))))