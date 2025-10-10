# NewCal Kotlin

A Kotlin library for working with ["The New Calendar"](https://thenewcalendar.com) - a reformed calendar system featuring 5 seasons and 9-day weeks.

## The New Calendar System

The New Calendar reimagines how we track time:

- **5 Seasons**: Winter, Spring, Summer, Autumn, and Fall - each exactly 73 days long
- **9-Day Weeks**: Named after the planets (Mercury, Venus, Earth, Mars, Jupiter, Saturn, Uranus, Neptune, Pluto)
- **Mid-Season Days**: Day 37 of each season is a special "Mid Season" day that doesn't belong to any week
- **Leap Days**: Leap years have a special "Leap Day" between Winter and Spring (occurring on day 70 of the year)
- **365/366 Days**: Standard years have 365 days (5 × 73), leap years add one extra day

### Calendar Structure

Each season has 73 days divided into:
- **Early Half**: Days 1-36 (4 complete weeks)
- **Mid-Season Day**: Day 37 (standalone day)
- **Late Half**: Days 38-73 (4 complete weeks)

## Installation

### Gradle (Kotlin DSL)

```kotlin
dependencies {
    implementation("com.ralphlandon.newcal:newcal-kotlin:1.0.0")
}
```

### Gradle (Groovy)

```groovy
dependencies {
    implementation 'com.ralphlandon.newcal:newcal-kotlin:1.0.0'
}
```

### Maven

```xml
<dependency>
    <groupId>com.ralphlandon.newcal</groupId>
    <artifactId>newcal-kotlin</artifactId>
    <version>1.0.0</version>
</dependency>
```

## Usage

### Convert Unix timestamp to New Calendar date

```kotlin
import com.ralphlandon.newcal.Date
import java.time.Instant

// From Unix timestamp (seconds)
val date = Date.fromUnix(1234567890)
println(date) // "Pluto, 55 Late Winter, 2009"

// From current time
val now = Date.fromUnix(Instant.now().epochSecond)
println(now)

// Access individual components
println("Year: ${date.year}")           // Year: 2009
println("Season: ${date.month}")        // Season: Winter
println("Day: ${date.day}")             // Day: 55
println("Day of week: ${date.getDayOfWeek()}")  // Day of week: Pluto
println("Month: ${date.getMonthWithModifier()}") // Month: Late Winter
```

### Check leap years

```kotlin
import com.ralphlandon.newcal.Date

val isLeap = Date.isLeapYear(2024)  // true
val isNotLeap = Date.isLeapYear(2023)  // false
```

### Working with leap days

```kotlin
import com.ralphlandon.newcal.Date
import java.time.LocalDateTime
import java.time.ZoneOffset

// February 29, 1972 is a leap day in the New Calendar
val leapDay = Date.fromUnix(
    LocalDateTime.of(1972, 2, 29, 0, 0, 0)
        .toEpochSecond(ZoneOffset.UTC)
)

println(leapDay)                    // "Leap Day, 1972"
println(leapDay.isLeapDay())        // true
println(leapDay.getDayOfWeek())     // "Leap Day"
```

### Mid-season days

```kotlin
import com.ralphlandon.newcal.Date

// Day 37 of any season is a mid-season day
val midSeasonDay = Date(2024, "Winter", 37)
println(midSeasonDay.getDayOfWeek())  // "Mid Season"
println(midSeasonDay.getMonthWithModifier())  // "Mid Winter"
println(midSeasonDay)  // "37 Mid Winter, 2024"
```

## Building

Build the library:

```bash
./gradlew build
```

Run tests:

```bash
./gradlew test
```

## License

MIT

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
