import org.jetbrains.kotlin.gradle.tasks.KotlinCompile

plugins {
    kotlin("jvm") version "1.6.10"
    application
}

group = "org.apprit"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
}

dependencies {
    implementation("io.ktor", "ktor-server-netty", getPropertyValue("ktor_version"))
    implementation("com.slack.api", "bolt", getPropertyValue("slack_sdk_version"))
    implementation("ch.qos.logback", "logback-classic", getPropertyValue("logback_version"))
    testImplementation(kotlin("test"))
    testImplementation("io.ktor", "ktor-server-tests", getPropertyValue("ktor_version"))
}

tasks.test {
    useJUnitPlatform()
}

tasks.withType<KotlinCompile> {
    kotlinOptions.jvmTarget = "11"
}

application {
    mainClass.set("ApplicationKT")
}

fun getPropertyValue(name: String): String {
    return project.properties[name].toString()
}