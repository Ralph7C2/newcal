plugins {
    kotlin("jvm") version "1.9.22"
    `maven-publish`
}

group = "com.ralphlandon.newcal"
version = "1.0.0"

repositories {
    mavenCentral()
}

dependencies {
    implementation(kotlin("stdlib"))
    testImplementation("org.junit.jupiter:junit-jupiter:5.10.1")
    testRuntimeOnly("org.junit.platform:junit-platform-launcher")
}

kotlin {
    jvmToolchain(11)
}

tasks.test {
    useJUnitPlatform()
}

publishing {
    publications {
        create<MavenPublication>("maven") {
            from(components["java"])

            pom {
                name.set("NewCal Kotlin")
                description.set("A Kotlin library for working with The New Calendar - a reformed calendar system featuring 5 seasons and 9-day weeks")
                url.set("https://github.com/ralph7c2/newcal")

                licenses {
                    license {
                        name.set("MIT License")
                        url.set("https://opensource.org/licenses/MIT")
                    }
                }
            }
        }
    }
}
