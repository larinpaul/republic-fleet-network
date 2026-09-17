package com.republic.hull.repositories

import com.republic.hull.models.HullStatus
import com.republic.hull.models.StarshipHull
import org.springframework.data.jpa.repository.JpaRepository
import org.springframework.stereotype.Repository
import java.util.UUID

/**
 * The data access layer for StarshipHull entities.
 * Spring Data JPA automatically implements this interface at runtime!
 */
@Repository
interface StarshipHullRepository : JpaRepository<StarshipHull, UUID> {

    /**
     * Custom query method.
     * Spring automatically translates this method name into:
     * SELECT * FROM starship_hulls WHERE status = ?
     */
    fun findByStatus(status: HullStatus): List<StarshipHull>

    /**
     * Another custom query.
     * Translates to: SELECT * FROM starship_hulls WHERE class_name = ?
     */
    fun findByClassName(className: String): List<StarshipHull>
}
