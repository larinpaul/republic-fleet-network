package com.republic.hull.controllers

import com.republic.hull.models.ShipBlueprintRequest
import com.republic.hull.models.ShipManifestDto
import com.republic.hull.services.HullFabricationService
import org.springframework.http.HttpStatus
import org.springframework.http.ResponseEntity
import org.springframework.web.bind.annotation.*
import java.util.UUID

/**
 * The Controller is the comms officer.
 * It listens for HTTP requests (transmissions)
 * from the galaxy, validates them, passes them to the Serivce (the foreman),
 * and sends the HTTP response back.
 *
 * The REST API entry point for the Hull Fabrication Service.
 * 
 * "All incoming transmissions are routed through this array."
 */
@RestController
@RequestMapping("/api/v1/ships")
@CrossOrigin(origins = ["*"]) // Allow the HoloNet (frontend/other services) to talk to us
class HullFabricationController(
    private val fabricationService: HullFabricationService
) {

    /**
     * POST /api/v1/ships/fabricate
     * Receives a blueprint and starts building the hull.
     */
    @PostMapping("/fabricate")
    fun fabricateStarfighter(@RequestBody blueprint: ShipBlueprintRequest): ResponseEntity<ShipManifestDto> {
        // The foreman does the work
        val manifest = fabricationService.fabricateShip(blueprint)
        
        // Return HTTP 201 Created with the manifest in the body
        return ResponseEntity.status(HttpStatus.CREATED).body(manifest)
    }

    /**
     * GET /api/v1/ships/{manifestId}
     * Checks the status of an existing hull.
     */
    @GetMapping("/{manifestId}")
    fun getShipStatus(@PathVariable manifestId: UUID): ResponseEntity<ShipManifestDto> {
        return try {
            val manifest = fabricationService.getShipStatus(manifestId)
            ResponseEntity.ok(manifest) // HTTP 200 OK
        } catch (e: NoSuchElementException) {
            ResponseEntity.notFound().build() // HTTP 404 Not Found
        }
    }
    
    /**
     * GET /api/v1/ships/health
     * A simple ping to check if the shipyard comms are online.
     */
    @GetMapping("/health")
    fun healthCheck(): ResponseEntity<Map<String, String>> {
        val status = mapOf(
            "status" to "ONLINE",
            "service" to "Hull Fabrication",
            "message" to "The Force is strong with this shipyard."
        )
        return ResponseEntity.ok(status)
    }
}
