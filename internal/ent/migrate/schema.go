// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountsColumns holds the columns for the "accounts" table.
	AccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 20},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "handle", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "bio", Type: field.TypeString, Nullable: true},
		{Name: "admin", Type: field.TypeBool, Default: false},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:       "accounts",
		Columns:    AccountsColumns,
		PrimaryKey: []*schema.Column{AccountsColumns[0]},
	}
	// AssetsColumns holds the columns for the "assets" table.
	AssetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "url", Type: field.TypeString},
		{Name: "mimetype", Type: field.TypeString},
		{Name: "width", Type: field.TypeInt},
		{Name: "height", Type: field.TypeInt},
		{Name: "account_id", Type: field.TypeString, Size: 20},
	}
	// AssetsTable holds the schema information for the "assets" table.
	AssetsTable = &schema.Table{
		Name:       "assets",
		Columns:    AssetsColumns,
		PrimaryKey: []*schema.Column{AssetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "assets_accounts_assets",
				Columns:    []*schema.Column{AssetsColumns[7]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// AuthenticationsColumns holds the columns for the "authentications" table.
	AuthenticationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 20},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "service", Type: field.TypeString},
		{Name: "identifier", Type: field.TypeString},
		{Name: "token", Type: field.TypeString},
		{Name: "metadata", Type: field.TypeJSON, Nullable: true},
		{Name: "account_authentication", Type: field.TypeString, Size: 20},
	}
	// AuthenticationsTable holds the schema information for the "authentications" table.
	AuthenticationsTable = &schema.Table{
		Name:       "authentications",
		Columns:    AuthenticationsColumns,
		PrimaryKey: []*schema.Column{AuthenticationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "authentications_accounts_authentication",
				Columns:    []*schema.Column{AuthenticationsColumns[6]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "authentication_service_identifier",
				Unique:  true,
				Columns: []*schema.Column{AuthenticationsColumns[2], AuthenticationsColumns[3]},
			},
		},
	}
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 20},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "slug", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString, Default: "(No description)"},
		{Name: "colour", Type: field.TypeString, Default: "#8577ce"},
		{Name: "sort", Type: field.TypeInt, Default: -1},
		{Name: "admin", Type: field.TypeBool, Default: false},
		{Name: "metadata", Type: field.TypeJSON, Nullable: true},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
	}
	// ClustersColumns holds the columns for the "clusters" table.
	ClustersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 20},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "slug", Type: field.TypeString, Unique: true},
		{Name: "image_url", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString},
		{Name: "properties", Type: field.TypeJSON, Nullable: true},
		{Name: "account_id", Type: field.TypeString, Size: 20},
		{Name: "parent_cluster_id", Type: field.TypeString, Nullable: true, Size: 20},
	}
	// ClustersTable holds the schema information for the "clusters" table.
	ClustersTable = &schema.Table{
		Name:       "clusters",
		Columns:    ClustersColumns,
		PrimaryKey: []*schema.Column{ClustersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "clusters_accounts_clusters",
				Columns:    []*schema.Column{ClustersColumns[9]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "clusters_clusters_clusters",
				Columns:    []*schema.Column{ClustersColumns[10]},
				RefColumns: []*schema.Column{ClustersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "cluster_slug",
				Unique:  false,
				Columns: []*schema.Column{ClustersColumns[5]},
			},
		},
	}
	// CollectionsColumns holds the columns for the "collections" table.
	CollectionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 20},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "account_collections", Type: field.TypeString, Nullable: true, Size: 20},
	}
	// CollectionsTable holds the schema information for the "collections" table.
	CollectionsTable = &schema.Table{
		Name:       "collections",
		Columns:    CollectionsColumns,
		PrimaryKey: []*schema.Column{CollectionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "collections_accounts_collections",
				Columns:    []*schema.Column{CollectionsColumns[5]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ItemsColumns holds the columns for the "items" table.
	ItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 20},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "slug", Type: field.TypeString},
		{Name: "image_url", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString},
		{Name: "properties", Type: field.TypeJSON, Nullable: true},
		{Name: "account_id", Type: field.TypeString, Size: 20},
	}
	// ItemsTable holds the schema information for the "items" table.
	ItemsTable = &schema.Table{
		Name:       "items",
		Columns:    ItemsColumns,
		PrimaryKey: []*schema.Column{ItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "items_accounts_items",
				Columns:    []*schema.Column{ItemsColumns[9]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// NotificationsColumns holds the columns for the "notifications" table.
	NotificationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 20},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "link", Type: field.TypeString},
		{Name: "read", Type: field.TypeBool},
	}
	// NotificationsTable holds the schema information for the "notifications" table.
	NotificationsTable = &schema.Table{
		Name:       "notifications",
		Columns:    NotificationsColumns,
		PrimaryKey: []*schema.Column{NotificationsColumns[0]},
	}
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 20},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "first", Type: field.TypeBool},
		{Name: "title", Type: field.TypeString, Nullable: true},
		{Name: "slug", Type: field.TypeString, Nullable: true},
		{Name: "pinned", Type: field.TypeBool, Default: false},
		{Name: "body", Type: field.TypeString},
		{Name: "short", Type: field.TypeString},
		{Name: "metadata", Type: field.TypeJSON, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"draft", "published"}, Default: "draft"},
		{Name: "url", Type: field.TypeString, Nullable: true},
		{Name: "account_posts", Type: field.TypeString, Size: 20},
		{Name: "category_id", Type: field.TypeString, Nullable: true, Size: 20},
		{Name: "root_post_id", Type: field.TypeString, Nullable: true, Size: 20},
		{Name: "reply_to_post_id", Type: field.TypeString, Nullable: true, Size: 20},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:       "posts",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "posts_accounts_posts",
				Columns:    []*schema.Column{PostsColumns[13]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "posts_categories_posts",
				Columns:    []*schema.Column{PostsColumns[14]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "posts_posts_posts",
				Columns:    []*schema.Column{PostsColumns[15]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "posts_posts_replies",
				Columns:    []*schema.Column{PostsColumns[16]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ReactsColumns holds the columns for the "reacts" table.
	ReactsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 20},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "emoji", Type: field.TypeString},
		{Name: "account_id", Type: field.TypeString, Size: 20},
		{Name: "post_id", Type: field.TypeString, Size: 20},
	}
	// ReactsTable holds the schema information for the "reacts" table.
	ReactsTable = &schema.Table{
		Name:       "reacts",
		Columns:    ReactsColumns,
		PrimaryKey: []*schema.Column{ReactsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "reacts_accounts_reacts",
				Columns:    []*schema.Column{ReactsColumns[3]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "reacts_posts_reacts",
				Columns:    []*schema.Column{ReactsColumns[4]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 20},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:       "roles",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
	}
	// SettingsColumns holds the columns for the "settings" table.
	SettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "value", Type: field.TypeString},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// SettingsTable holds the schema information for the "settings" table.
	SettingsTable = &schema.Table{
		Name:       "settings",
		Columns:    SettingsColumns,
		PrimaryKey: []*schema.Column{SettingsColumns[0]},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 20},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
	}
	// AccountTagsColumns holds the columns for the "account_tags" table.
	AccountTagsColumns = []*schema.Column{
		{Name: "account_id", Type: field.TypeString, Size: 20},
		{Name: "tag_id", Type: field.TypeString, Size: 20},
	}
	// AccountTagsTable holds the schema information for the "account_tags" table.
	AccountTagsTable = &schema.Table{
		Name:       "account_tags",
		Columns:    AccountTagsColumns,
		PrimaryKey: []*schema.Column{AccountTagsColumns[0], AccountTagsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "account_tags_account_id",
				Columns:    []*schema.Column{AccountTagsColumns[0]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "account_tags_tag_id",
				Columns:    []*schema.Column{AccountTagsColumns[1]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ClusterItemsColumns holds the columns for the "cluster_items" table.
	ClusterItemsColumns = []*schema.Column{
		{Name: "cluster_id", Type: field.TypeString, Size: 20},
		{Name: "item_id", Type: field.TypeString, Size: 20},
	}
	// ClusterItemsTable holds the schema information for the "cluster_items" table.
	ClusterItemsTable = &schema.Table{
		Name:       "cluster_items",
		Columns:    ClusterItemsColumns,
		PrimaryKey: []*schema.Column{ClusterItemsColumns[0], ClusterItemsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cluster_items_cluster_id",
				Columns:    []*schema.Column{ClusterItemsColumns[0]},
				RefColumns: []*schema.Column{ClustersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "cluster_items_item_id",
				Columns:    []*schema.Column{ClusterItemsColumns[1]},
				RefColumns: []*schema.Column{ItemsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ClusterAssetsColumns holds the columns for the "cluster_assets" table.
	ClusterAssetsColumns = []*schema.Column{
		{Name: "cluster_id", Type: field.TypeString, Size: 20},
		{Name: "asset_id", Type: field.TypeString},
	}
	// ClusterAssetsTable holds the schema information for the "cluster_assets" table.
	ClusterAssetsTable = &schema.Table{
		Name:       "cluster_assets",
		Columns:    ClusterAssetsColumns,
		PrimaryKey: []*schema.Column{ClusterAssetsColumns[0], ClusterAssetsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cluster_assets_cluster_id",
				Columns:    []*schema.Column{ClusterAssetsColumns[0]},
				RefColumns: []*schema.Column{ClustersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "cluster_assets_asset_id",
				Columns:    []*schema.Column{ClusterAssetsColumns[1]},
				RefColumns: []*schema.Column{AssetsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// CollectionPostsColumns holds the columns for the "collection_posts" table.
	CollectionPostsColumns = []*schema.Column{
		{Name: "collection_id", Type: field.TypeString, Size: 20},
		{Name: "post_id", Type: field.TypeString, Size: 20},
	}
	// CollectionPostsTable holds the schema information for the "collection_posts" table.
	CollectionPostsTable = &schema.Table{
		Name:       "collection_posts",
		Columns:    CollectionPostsColumns,
		PrimaryKey: []*schema.Column{CollectionPostsColumns[0], CollectionPostsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "collection_posts_collection_id",
				Columns:    []*schema.Column{CollectionPostsColumns[0]},
				RefColumns: []*schema.Column{CollectionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "collection_posts_post_id",
				Columns:    []*schema.Column{CollectionPostsColumns[1]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ItemAssetsColumns holds the columns for the "item_assets" table.
	ItemAssetsColumns = []*schema.Column{
		{Name: "item_id", Type: field.TypeString, Size: 20},
		{Name: "asset_id", Type: field.TypeString},
	}
	// ItemAssetsTable holds the schema information for the "item_assets" table.
	ItemAssetsTable = &schema.Table{
		Name:       "item_assets",
		Columns:    ItemAssetsColumns,
		PrimaryKey: []*schema.Column{ItemAssetsColumns[0], ItemAssetsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "item_assets_item_id",
				Columns:    []*schema.Column{ItemAssetsColumns[0]},
				RefColumns: []*schema.Column{ItemsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "item_assets_asset_id",
				Columns:    []*schema.Column{ItemAssetsColumns[1]},
				RefColumns: []*schema.Column{AssetsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PostAssetsColumns holds the columns for the "post_assets" table.
	PostAssetsColumns = []*schema.Column{
		{Name: "post_id", Type: field.TypeString, Size: 20},
		{Name: "asset_id", Type: field.TypeString},
	}
	// PostAssetsTable holds the schema information for the "post_assets" table.
	PostAssetsTable = &schema.Table{
		Name:       "post_assets",
		Columns:    PostAssetsColumns,
		PrimaryKey: []*schema.Column{PostAssetsColumns[0], PostAssetsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "post_assets_post_id",
				Columns:    []*schema.Column{PostAssetsColumns[0]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "post_assets_asset_id",
				Columns:    []*schema.Column{PostAssetsColumns[1]},
				RefColumns: []*schema.Column{AssetsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// RoleAccountsColumns holds the columns for the "role_accounts" table.
	RoleAccountsColumns = []*schema.Column{
		{Name: "role_id", Type: field.TypeString, Size: 20},
		{Name: "account_id", Type: field.TypeString, Size: 20},
	}
	// RoleAccountsTable holds the schema information for the "role_accounts" table.
	RoleAccountsTable = &schema.Table{
		Name:       "role_accounts",
		Columns:    RoleAccountsColumns,
		PrimaryKey: []*schema.Column{RoleAccountsColumns[0], RoleAccountsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "role_accounts_role_id",
				Columns:    []*schema.Column{RoleAccountsColumns[0]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "role_accounts_account_id",
				Columns:    []*schema.Column{RoleAccountsColumns[1]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TagPostsColumns holds the columns for the "tag_posts" table.
	TagPostsColumns = []*schema.Column{
		{Name: "tag_id", Type: field.TypeString, Size: 20},
		{Name: "post_id", Type: field.TypeString, Size: 20},
	}
	// TagPostsTable holds the schema information for the "tag_posts" table.
	TagPostsTable = &schema.Table{
		Name:       "tag_posts",
		Columns:    TagPostsColumns,
		PrimaryKey: []*schema.Column{TagPostsColumns[0], TagPostsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tag_posts_tag_id",
				Columns:    []*schema.Column{TagPostsColumns[0]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "tag_posts_post_id",
				Columns:    []*schema.Column{TagPostsColumns[1]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TagClustersColumns holds the columns for the "tag_clusters" table.
	TagClustersColumns = []*schema.Column{
		{Name: "tag_id", Type: field.TypeString, Size: 20},
		{Name: "cluster_id", Type: field.TypeString, Size: 20},
	}
	// TagClustersTable holds the schema information for the "tag_clusters" table.
	TagClustersTable = &schema.Table{
		Name:       "tag_clusters",
		Columns:    TagClustersColumns,
		PrimaryKey: []*schema.Column{TagClustersColumns[0], TagClustersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tag_clusters_tag_id",
				Columns:    []*schema.Column{TagClustersColumns[0]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "tag_clusters_cluster_id",
				Columns:    []*schema.Column{TagClustersColumns[1]},
				RefColumns: []*schema.Column{ClustersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TagItemsColumns holds the columns for the "tag_items" table.
	TagItemsColumns = []*schema.Column{
		{Name: "tag_id", Type: field.TypeString, Size: 20},
		{Name: "item_id", Type: field.TypeString, Size: 20},
	}
	// TagItemsTable holds the schema information for the "tag_items" table.
	TagItemsTable = &schema.Table{
		Name:       "tag_items",
		Columns:    TagItemsColumns,
		PrimaryKey: []*schema.Column{TagItemsColumns[0], TagItemsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tag_items_tag_id",
				Columns:    []*schema.Column{TagItemsColumns[0]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "tag_items_item_id",
				Columns:    []*schema.Column{TagItemsColumns[1]},
				RefColumns: []*schema.Column{ItemsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountsTable,
		AssetsTable,
		AuthenticationsTable,
		CategoriesTable,
		ClustersTable,
		CollectionsTable,
		ItemsTable,
		NotificationsTable,
		PostsTable,
		ReactsTable,
		RolesTable,
		SettingsTable,
		TagsTable,
		AccountTagsTable,
		ClusterItemsTable,
		ClusterAssetsTable,
		CollectionPostsTable,
		ItemAssetsTable,
		PostAssetsTable,
		RoleAccountsTable,
		TagPostsTable,
		TagClustersTable,
		TagItemsTable,
	}
)

func init() {
	AssetsTable.ForeignKeys[0].RefTable = AccountsTable
	AuthenticationsTable.ForeignKeys[0].RefTable = AccountsTable
	ClustersTable.ForeignKeys[0].RefTable = AccountsTable
	ClustersTable.ForeignKeys[1].RefTable = ClustersTable
	CollectionsTable.ForeignKeys[0].RefTable = AccountsTable
	ItemsTable.ForeignKeys[0].RefTable = AccountsTable
	PostsTable.ForeignKeys[0].RefTable = AccountsTable
	PostsTable.ForeignKeys[1].RefTable = CategoriesTable
	PostsTable.ForeignKeys[2].RefTable = PostsTable
	PostsTable.ForeignKeys[3].RefTable = PostsTable
	ReactsTable.ForeignKeys[0].RefTable = AccountsTable
	ReactsTable.ForeignKeys[1].RefTable = PostsTable
	AccountTagsTable.ForeignKeys[0].RefTable = AccountsTable
	AccountTagsTable.ForeignKeys[1].RefTable = TagsTable
	ClusterItemsTable.ForeignKeys[0].RefTable = ClustersTable
	ClusterItemsTable.ForeignKeys[1].RefTable = ItemsTable
	ClusterAssetsTable.ForeignKeys[0].RefTable = ClustersTable
	ClusterAssetsTable.ForeignKeys[1].RefTable = AssetsTable
	CollectionPostsTable.ForeignKeys[0].RefTable = CollectionsTable
	CollectionPostsTable.ForeignKeys[1].RefTable = PostsTable
	ItemAssetsTable.ForeignKeys[0].RefTable = ItemsTable
	ItemAssetsTable.ForeignKeys[1].RefTable = AssetsTable
	PostAssetsTable.ForeignKeys[0].RefTable = PostsTable
	PostAssetsTable.ForeignKeys[1].RefTable = AssetsTable
	RoleAccountsTable.ForeignKeys[0].RefTable = RolesTable
	RoleAccountsTable.ForeignKeys[1].RefTable = AccountsTable
	TagPostsTable.ForeignKeys[0].RefTable = TagsTable
	TagPostsTable.ForeignKeys[1].RefTable = PostsTable
	TagClustersTable.ForeignKeys[0].RefTable = TagsTable
	TagClustersTable.ForeignKeys[1].RefTable = ClustersTable
	TagItemsTable.ForeignKeys[0].RefTable = TagsTable
	TagItemsTable.ForeignKeys[1].RefTable = ItemsTable
}
