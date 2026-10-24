class MyCalendarTwo
    def initialize()
        @single_booked = []
        @double_booked = []
    end

    def book(start_time, end_time)
        new_event = [start_time, end_time]
        @double_booked.each do |event|
            return false if intersection?(event, new_event)
        end

        @single_booked.each do |event|
            next unless intersection?(event, new_event)
            
            @double_booked << [
                [event[0], new_event[0]].max,
                [event[1], new_event[1]].min
            ]
        end

        @single_booked << new_event
        true
    end

    private

        def intersection?(event1, event2)
           event1[1] > event2[0] && event2[1] > event1[0] 
        end
end