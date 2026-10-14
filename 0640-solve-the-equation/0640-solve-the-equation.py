class Solution:
    def solveEquation(self, equation: str) -> str:
        def parse(expr):
            expr = expr.replace('-', '+-')  # Make every term explicit with '+'
            tokens = expr.split('+')
            x_coeff = 0
            const = 0
            for token in tokens:
                if not token:
                    continue
                if 'x' in token:
                    if token == 'x':
                        x_coeff += 1
                    elif token == '-x':
                        x_coeff -= 1
                    else:
                        x_coeff += int(token.replace('x', ''))
                else:
                    const += int(token)
            return x_coeff, const

        left, right = equation.split('=')
        left_x, left_const = parse(left)
        right_x, right_const = parse(right)

        # Move x terms to one side and constants to the other
        total_x = left_x - right_x
        total_const = right_const - left_const

        if total_x == 0:
            if total_const == 0:
                return "Infinite solutions"
            else:
                return "No solution"
        else:
            x_val = total_const // total_x
            return f"x={x_val}"