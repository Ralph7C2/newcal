package com.ralphlandon.newcal

import org.junit.jupiter.api.Test
import org.junit.jupiter.api.Assertions.*
import java.time.LocalDateTime
import java.time.ZoneOffset

class NewCalTest {

    private fun toUnix(year: Int, month: Int, day: Int): Long {
        return LocalDateTime.of(year, month, day, 0, 0, 0)
            .toEpochSecond(ZoneOffset.UTC)
    }

    @Test
    fun testFromUnix() {
        // Tests from the Go implementation
        assertEquals(Date(1968, "Autumn", 73), Date.fromUnix(toUnix(1968, 10, 8)))
        assertEquals(Date(1968, "Fall", 1), Date.fromUnix(toUnix(1968, 10, 9)))
        assertEquals(Date(1968, "Fall", 73), Date.fromUnix(toUnix(1968, 12, 20)))
        assertEquals(Date(1969, "Winter", 1), Date.fromUnix(toUnix(1968, 12, 21)))
        assertEquals(Date(1969, "Fall", 73), Date.fromUnix(toUnix(1969, 12, 20)))
        assertEquals(Date(1970, "Winter", 1), Date.fromUnix(toUnix(1969, 12, 21)))
        assertEquals(Date(1970, "Winter", 11), Date.fromUnix(-86400))
        assertEquals(Date(1970, "Winter", 11), Date.fromUnix(toUnix(1969, 12, 31)))
        assertEquals(Date(1970, "Winter", 12), Date.fromUnix(0))
        assertEquals(Date(1970, "Winter", 12), Date.fromUnix(toUnix(1970, 1, 1)))
        assertEquals(Date(1970, "Winter", 13), Date.fromUnix(86400))
        assertEquals(Date(1970, "Winter", 13), Date.fromUnix(toUnix(1970, 1, 2)))
        assertEquals(Date(1970, "Winter", 71), Date.fromUnix(5097600))
        assertEquals(Date(1970, "Winter", 71), Date.fromUnix(toUnix(1970, 3, 1)))
        assertEquals(Date(1970, "Winter", 73), Date.fromUnix(toUnix(1970, 3, 3)))
        assertEquals(Date(1970, "Spring", 1), Date.fromUnix(toUnix(1970, 3, 4)))
        assertEquals(Date(1971, "Winter", 71), Date.fromUnix(36633600))
        assertEquals(Date(1971, "Winter", 71), Date.fromUnix(toUnix(1971, 3, 1)))
        assertEquals(Date(1972, "Leap Day", 0), Date.fromUnix(36633600 + 365 * 86400L))
        assertEquals(Date(1972, "Leap Day", 0), Date.fromUnix(toUnix(1972, 2, 29)))
        assertEquals(Date(1972, "Winter", 71), Date.fromUnix(36633600 + 366 * 86400L))
        assertEquals(Date(1987, "Spring", 73), Date.fromUnix(toUnix(1987, 5, 15)))
        assertEquals(Date(1999, "Summer", 1), Date.fromUnix(toUnix(1999, 5, 16)))
        assertEquals(Date(2999, "Summer", 73), Date.fromUnix(toUnix(2999, 7, 27)))
        assertEquals(Date(3147, "Autumn", 1), Date.fromUnix(toUnix(3147, 7, 28)))
        assertEquals(Date(3147, "Autumn", 73), Date.fromUnix(toUnix(3147, 10, 8)))
        assertEquals(Date(1, "Autumn", 73), Date.fromUnix(toUnix(1, 10, 8)))
        assertEquals(Date(0, "Autumn", 73), Date.fromUnix(toUnix(0, 10, 8)))
    }

    @Test
    fun testDayOfWeek() {
        val tests = listOf(
            0 to "Leap Day",
            1 to "Mercury",
            2 to "Venus",
            3 to "Earth",
            4 to "Mars",
            5 to "Jupiter",
            6 to "Saturn",
            7 to "Uranus",
            8 to "Neptune",
            9 to "Pluto",
            10 to "Mercury",
            11 to "Venus",
            29 to "Venus",
            36 to "Pluto",
            37 to "Mid Season",
            38 to "Mercury"
        )

        for ((day, expected) in tests) {
            val date = Date(2000, "Winter", day)
            assertEquals(expected, date.getDayOfWeek(), "For day $day")
        }
    }

    @Test
    fun testMonthWithModifier() {
        assertEquals("Early Winter", Date(2000, "Winter", 1).getMonthWithModifier())
        assertEquals("Early Winter", Date(2000, "Winter", 36).getMonthWithModifier())
        assertEquals("Mid Winter", Date(2000, "Winter", 37).getMonthWithModifier())
        assertEquals("Late Winter", Date(2000, "Winter", 38).getMonthWithModifier())
        assertEquals("Late Winter", Date(2000, "Winter", 73).getMonthWithModifier())
    }

    @Test
    fun testIsLeapDay() {
        assertTrue(Date(1972, "Leap Day", -1).isLeapDay())
        assertFalse(Date(2000, "Winter", 1).isLeapDay())
    }

    @Test
    fun testToString() {
        assertEquals("Mercury, 1 Early Winter, 1970", Date(1970, "Winter", 1).toString())
        assertEquals("37 Mid Winter, 1970", Date(1970, "Winter", 37).toString())
        assertEquals("Mercury, 38 Late Winter, 1970", Date(1970, "Winter", 38).toString())
        assertEquals("Leap Day, 1972", Date(1972, "Leap Day", -1).toString())
    }

    @Test
    fun testIsLeapYear() {
        assertTrue(Date.isLeapYear(2000))
        assertTrue(Date.isLeapYear(2004))
        assertTrue(Date.isLeapYear(1972))
        assertFalse(Date.isLeapYear(1900))
        assertFalse(Date.isLeapYear(2001))
        assertFalse(Date.isLeapYear(1970))
    }
}
