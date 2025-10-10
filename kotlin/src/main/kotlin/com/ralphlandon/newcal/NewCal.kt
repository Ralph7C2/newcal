package com.ralphlandon.newcal

/**
 * Constants for seasons
 */
object Season {
    const val WINTER = "Winter"
    const val SPRING = "Spring"
    const val SUMMER = "Summer"
    const val AUTUMN = "Autumn"
    const val FALL = "Fall"

    val all = listOf(WINTER, SPRING, SUMMER, AUTUMN, FALL)
}

/**
 * Constants for days of the week
 */
object DayOfWeek {
    const val MID_SEASON = "Mid Season"
    const val MERCURY = "Mercury"
    const val VENUS = "Venus"
    const val EARTH = "Earth"
    const val MARS = "Mars"
    const val JUPITER = "Jupiter"
    const val SATURN = "Saturn"
    const val URANUS = "Uranus"
    const val NEPTUNE = "Neptune"
    const val PLUTO = "Pluto"

    val planets = listOf(MERCURY, VENUS, EARTH, MARS, JUPITER, SATURN, URANUS, NEPTUNE, PLUTO)
}

/**
 * Represents a date in the New Calendar system
 *
 * @property year The year
 * @property month The season name
 * @property day The day of the season (1-73), or -1 for Leap Day
 */
data class Date(
    val year: Int,
    val month: String,
    val day: Int
) {
    companion object {
        const val LEAP_DAY = -1

        /**
         * Converts a Unix timestamp to a New Calendar date
         *
         * @param unix The Unix timestamp in seconds
         * @return The corresponding New Calendar date
         */
        fun fromUnix(unix: Long): Date {
            var days = unix / 86400
            days += 11
            var year = 1970

            // Handle dates before Unix epoch
            while (days < 0) {
                year--
                days += if (isLeapYear(year)) 366 else 365
            }

            // Handle dates after Unix epoch
            while (days > 365) {
                if (isLeapYear(year)) {
                    if (days >= 366) {
                        days -= 366
                        year++
                    } else {
                        break
                    }
                } else {
                    days -= 365
                    year++
                }
            }

            // Handle leap day
            if (isLeapYear(year)) {
                if (days == 70L) {
                    return Date(
                        year = year,
                        month = "Leap Day",
                        day = LEAP_DAY
                    )
                }
                if (days > 70) {
                    days--
                }
            }

            val season = (days / 73).toInt()
            val dayOfSeason = (days % 73).toInt()

            return Date(
                year = year,
                month = Season.all[season],
                day = dayOfSeason + 1
            )
        }

        /**
         * Checks if a year is a leap year
         *
         * @param year The year to check
         * @return true if the year is a leap year, false otherwise
         */
        fun isLeapYear(year: Int): Boolean {
            return year % 4 == 0 && (year % 100 != 0 || year % 400 == 0)
        }
    }

    /**
     * Checks if this date is a leap day
     */
    fun isLeapDay(): Boolean = day == LEAP_DAY

    /**
     * Gets the day of the week for this date
     */
    fun getDayOfWeek(): String {
        if (day == LEAP_DAY) {
            return "Leap Day"
        }
        if (day == 37) {
            return DayOfWeek.MID_SEASON
        }

        var adjustedDay = day
        if (day > 37) {
            adjustedDay--
        }

        val dayIndex = (adjustedDay - 1) % DayOfWeek.planets.size
        return DayOfWeek.planets[dayIndex]
    }

    /**
     * Gets the month name with modifier (Early/Mid/Late)
     */
    fun getMonthWithModifier(): String {
        val modifier = when {
            day > 37 -> "Late"
            day == 37 -> "Mid"
            else -> "Early"
        }
        return "$modifier $month"
    }

    /**
     * Converts the date to a human-readable string
     */
    override fun toString(): String {
        if (isLeapDay()) {
            return "Leap Day, $year"
        }

        val dayOfWeek = if (day != 37) {
            "${getDayOfWeek()}, "
        } else {
            ""
        }

        return "$dayOfWeek$day ${getMonthWithModifier()}, $year"
    }
}
