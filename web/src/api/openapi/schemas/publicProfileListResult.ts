/**
 * Generated by orval v6.17.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import type { PublicProfileList } from "./publicProfileList";

export interface PublicProfileListResult {
  page_size: number;
  results: number;
  total_pages: number;
  current_page: number;
  next_page?: number;
  profiles: PublicProfileList;
}
