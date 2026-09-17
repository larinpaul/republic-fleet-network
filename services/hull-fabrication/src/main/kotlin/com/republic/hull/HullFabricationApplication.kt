package com.republic.hull

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication

/**
 * The Republic Hull Fabrication Service.
 *
 * "A ship is only as strong as its hull, and a fleet
 *  is only as resilient as its supply chain."
 *  — Admiral Dodonna, Republic Fleet Command
 *
 * This is the main entry point for the Spring Boot application.
 * The @SpringBootApplication annotation magically wires together:
 *   - The web server (Tomcat)
 *   - The database connection (PostgreSQL via JPA)
 *   - The Kafka producer (HoloNet event bus)
 *   - All our controllers, services, and repositories
 */
@SpringBootApplication
class HullFabricationApplication

fun main(args: Array<String>) {
    runApplication<HullFabricationApplication>(*args)
}
