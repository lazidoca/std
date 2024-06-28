/**
 * Generated by orval v6.30.2 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import type { AssetIDs } from "./assetIDs";
import type { NodeName } from "./nodeName";
import type { PostContent } from "./postContent";
import type { Properties } from "./properties";
import type { Slug } from "./slug";
import type { Url } from "./url";

/**
 * Note: Properties are replace-all and are not merged with existing.

 */
export interface NodeMutableProps {
  asset_ids?: AssetIDs;
  content?: PostContent;
  name?: NodeName;
  parent?: Slug;
  properties?: Properties;
  slug?: Slug;
  url?: Url;
}
