/**
 * Generated by orval v6.9.6 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import type { AccountHandle } from "./accountHandle";
import type { TagListIDs } from "./tagListIDs";
import type { CategoryNameList } from "./categoryNameList";

export type ThreadListParams = {
  author?: AccountHandle;
  tags?: TagListIDs;
  categories?: CategoryNameList;
};