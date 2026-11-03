(require data/skip-list)

(define count-intervals%
  (class object%
    (super-new)

    (init-field [m (make-skip-list)]
                [cnt 0])

    (define (len lower upper)
      (- upper lower -1))

    (define/public (add left right)
      (let loop ([pos (dict-iterate-least/>=? m left)])
        (when pos
          (let ([lower (dict-iterate-value m pos)]
                [upper (dict-iterate-key m pos)])
            (when (<= lower right)
              (set! left (min left lower))
              (set! right (max right upper))
              (set! cnt (- cnt (len lower upper))) 

              (dict-remove! m upper)
              (loop (dict-iterate-next m pos))))))

      (set! cnt (+ cnt (len left right)))
      (dict-set! m right left))

    (define/public (count)
      cnt)))