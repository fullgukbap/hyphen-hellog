// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"hyphen-hellog/ent/migrate"

	"hyphen-hellog/ent/author"
	"hyphen-hellog/ent/comment"
	"hyphen-hellog/ent/post"
	"hyphen-hellog/ent/vote"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Author is the client for interacting with the Author builders.
	Author *AuthorClient
	// Comment is the client for interacting with the Comment builders.
	Comment *CommentClient
	// Post is the client for interacting with the Post builders.
	Post *PostClient
	// Vote is the client for interacting with the Vote builders.
	Vote *VoteClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Author = NewAuthorClient(c.config)
	c.Comment = NewCommentClient(c.config)
	c.Post = NewPostClient(c.config)
	c.Vote = NewVoteClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Author:  NewAuthorClient(cfg),
		Comment: NewCommentClient(cfg),
		Post:    NewPostClient(cfg),
		Vote:    NewVoteClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Author:  NewAuthorClient(cfg),
		Comment: NewCommentClient(cfg),
		Post:    NewPostClient(cfg),
		Vote:    NewVoteClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Author.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Author.Use(hooks...)
	c.Comment.Use(hooks...)
	c.Post.Use(hooks...)
	c.Vote.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Author.Intercept(interceptors...)
	c.Comment.Intercept(interceptors...)
	c.Post.Intercept(interceptors...)
	c.Vote.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *AuthorMutation:
		return c.Author.mutate(ctx, m)
	case *CommentMutation:
		return c.Comment.mutate(ctx, m)
	case *PostMutation:
		return c.Post.mutate(ctx, m)
	case *VoteMutation:
		return c.Vote.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// AuthorClient is a client for the Author schema.
type AuthorClient struct {
	config
}

// NewAuthorClient returns a client for the Author from the given config.
func NewAuthorClient(c config) *AuthorClient {
	return &AuthorClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `author.Hooks(f(g(h())))`.
func (c *AuthorClient) Use(hooks ...Hook) {
	c.hooks.Author = append(c.hooks.Author, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `author.Intercept(f(g(h())))`.
func (c *AuthorClient) Intercept(interceptors ...Interceptor) {
	c.inters.Author = append(c.inters.Author, interceptors...)
}

// Create returns a builder for creating a Author entity.
func (c *AuthorClient) Create() *AuthorCreate {
	mutation := newAuthorMutation(c.config, OpCreate)
	return &AuthorCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Author entities.
func (c *AuthorClient) CreateBulk(builders ...*AuthorCreate) *AuthorCreateBulk {
	return &AuthorCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *AuthorClient) MapCreateBulk(slice any, setFunc func(*AuthorCreate, int)) *AuthorCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &AuthorCreateBulk{err: fmt.Errorf("calling to AuthorClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*AuthorCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &AuthorCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Author.
func (c *AuthorClient) Update() *AuthorUpdate {
	mutation := newAuthorMutation(c.config, OpUpdate)
	return &AuthorUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AuthorClient) UpdateOne(a *Author) *AuthorUpdateOne {
	mutation := newAuthorMutation(c.config, OpUpdateOne, withAuthor(a))
	return &AuthorUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AuthorClient) UpdateOneID(id int) *AuthorUpdateOne {
	mutation := newAuthorMutation(c.config, OpUpdateOne, withAuthorID(id))
	return &AuthorUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Author.
func (c *AuthorClient) Delete() *AuthorDelete {
	mutation := newAuthorMutation(c.config, OpDelete)
	return &AuthorDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AuthorClient) DeleteOne(a *Author) *AuthorDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AuthorClient) DeleteOneID(id int) *AuthorDeleteOne {
	builder := c.Delete().Where(author.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AuthorDeleteOne{builder}
}

// Query returns a query builder for Author.
func (c *AuthorClient) Query() *AuthorQuery {
	return &AuthorQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeAuthor},
		inters: c.Interceptors(),
	}
}

// Get returns a Author entity by its id.
func (c *AuthorClient) Get(ctx context.Context, id int) (*Author, error) {
	return c.Query().Where(author.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AuthorClient) GetX(ctx context.Context, id int) *Author {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AuthorClient) Hooks() []Hook {
	return c.hooks.Author
}

// Interceptors returns the client interceptors.
func (c *AuthorClient) Interceptors() []Interceptor {
	return c.inters.Author
}

func (c *AuthorClient) mutate(ctx context.Context, m *AuthorMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&AuthorCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&AuthorUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&AuthorUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&AuthorDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Author mutation op: %q", m.Op())
	}
}

// CommentClient is a client for the Comment schema.
type CommentClient struct {
	config
}

// NewCommentClient returns a client for the Comment from the given config.
func NewCommentClient(c config) *CommentClient {
	return &CommentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `comment.Hooks(f(g(h())))`.
func (c *CommentClient) Use(hooks ...Hook) {
	c.hooks.Comment = append(c.hooks.Comment, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `comment.Intercept(f(g(h())))`.
func (c *CommentClient) Intercept(interceptors ...Interceptor) {
	c.inters.Comment = append(c.inters.Comment, interceptors...)
}

// Create returns a builder for creating a Comment entity.
func (c *CommentClient) Create() *CommentCreate {
	mutation := newCommentMutation(c.config, OpCreate)
	return &CommentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Comment entities.
func (c *CommentClient) CreateBulk(builders ...*CommentCreate) *CommentCreateBulk {
	return &CommentCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CommentClient) MapCreateBulk(slice any, setFunc func(*CommentCreate, int)) *CommentCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CommentCreateBulk{err: fmt.Errorf("calling to CommentClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CommentCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CommentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Comment.
func (c *CommentClient) Update() *CommentUpdate {
	mutation := newCommentMutation(c.config, OpUpdate)
	return &CommentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CommentClient) UpdateOne(co *Comment) *CommentUpdateOne {
	mutation := newCommentMutation(c.config, OpUpdateOne, withComment(co))
	return &CommentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CommentClient) UpdateOneID(id int) *CommentUpdateOne {
	mutation := newCommentMutation(c.config, OpUpdateOne, withCommentID(id))
	return &CommentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Comment.
func (c *CommentClient) Delete() *CommentDelete {
	mutation := newCommentMutation(c.config, OpDelete)
	return &CommentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CommentClient) DeleteOne(co *Comment) *CommentDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CommentClient) DeleteOneID(id int) *CommentDeleteOne {
	builder := c.Delete().Where(comment.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CommentDeleteOne{builder}
}

// Query returns a query builder for Comment.
func (c *CommentClient) Query() *CommentQuery {
	return &CommentQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeComment},
		inters: c.Interceptors(),
	}
}

// Get returns a Comment entity by its id.
func (c *CommentClient) Get(ctx context.Context, id int) (*Comment, error) {
	return c.Query().Where(comment.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CommentClient) GetX(ctx context.Context, id int) *Comment {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CommentClient) Hooks() []Hook {
	return c.hooks.Comment
}

// Interceptors returns the client interceptors.
func (c *CommentClient) Interceptors() []Interceptor {
	return c.inters.Comment
}

func (c *CommentClient) mutate(ctx context.Context, m *CommentMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CommentCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CommentUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CommentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CommentDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Comment mutation op: %q", m.Op())
	}
}

// PostClient is a client for the Post schema.
type PostClient struct {
	config
}

// NewPostClient returns a client for the Post from the given config.
func NewPostClient(c config) *PostClient {
	return &PostClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `post.Hooks(f(g(h())))`.
func (c *PostClient) Use(hooks ...Hook) {
	c.hooks.Post = append(c.hooks.Post, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `post.Intercept(f(g(h())))`.
func (c *PostClient) Intercept(interceptors ...Interceptor) {
	c.inters.Post = append(c.inters.Post, interceptors...)
}

// Create returns a builder for creating a Post entity.
func (c *PostClient) Create() *PostCreate {
	mutation := newPostMutation(c.config, OpCreate)
	return &PostCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Post entities.
func (c *PostClient) CreateBulk(builders ...*PostCreate) *PostCreateBulk {
	return &PostCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *PostClient) MapCreateBulk(slice any, setFunc func(*PostCreate, int)) *PostCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &PostCreateBulk{err: fmt.Errorf("calling to PostClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*PostCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &PostCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Post.
func (c *PostClient) Update() *PostUpdate {
	mutation := newPostMutation(c.config, OpUpdate)
	return &PostUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PostClient) UpdateOne(po *Post) *PostUpdateOne {
	mutation := newPostMutation(c.config, OpUpdateOne, withPost(po))
	return &PostUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PostClient) UpdateOneID(id int) *PostUpdateOne {
	mutation := newPostMutation(c.config, OpUpdateOne, withPostID(id))
	return &PostUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Post.
func (c *PostClient) Delete() *PostDelete {
	mutation := newPostMutation(c.config, OpDelete)
	return &PostDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PostClient) DeleteOne(po *Post) *PostDeleteOne {
	return c.DeleteOneID(po.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PostClient) DeleteOneID(id int) *PostDeleteOne {
	builder := c.Delete().Where(post.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PostDeleteOne{builder}
}

// Query returns a query builder for Post.
func (c *PostClient) Query() *PostQuery {
	return &PostQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePost},
		inters: c.Interceptors(),
	}
}

// Get returns a Post entity by its id.
func (c *PostClient) Get(ctx context.Context, id int) (*Post, error) {
	return c.Query().Where(post.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PostClient) GetX(ctx context.Context, id int) *Post {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PostClient) Hooks() []Hook {
	return c.hooks.Post
}

// Interceptors returns the client interceptors.
func (c *PostClient) Interceptors() []Interceptor {
	return c.inters.Post
}

func (c *PostClient) mutate(ctx context.Context, m *PostMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PostCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PostUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PostUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PostDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Post mutation op: %q", m.Op())
	}
}

// VoteClient is a client for the Vote schema.
type VoteClient struct {
	config
}

// NewVoteClient returns a client for the Vote from the given config.
func NewVoteClient(c config) *VoteClient {
	return &VoteClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `vote.Hooks(f(g(h())))`.
func (c *VoteClient) Use(hooks ...Hook) {
	c.hooks.Vote = append(c.hooks.Vote, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `vote.Intercept(f(g(h())))`.
func (c *VoteClient) Intercept(interceptors ...Interceptor) {
	c.inters.Vote = append(c.inters.Vote, interceptors...)
}

// Create returns a builder for creating a Vote entity.
func (c *VoteClient) Create() *VoteCreate {
	mutation := newVoteMutation(c.config, OpCreate)
	return &VoteCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Vote entities.
func (c *VoteClient) CreateBulk(builders ...*VoteCreate) *VoteCreateBulk {
	return &VoteCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *VoteClient) MapCreateBulk(slice any, setFunc func(*VoteCreate, int)) *VoteCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &VoteCreateBulk{err: fmt.Errorf("calling to VoteClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*VoteCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &VoteCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Vote.
func (c *VoteClient) Update() *VoteUpdate {
	mutation := newVoteMutation(c.config, OpUpdate)
	return &VoteUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *VoteClient) UpdateOne(v *Vote) *VoteUpdateOne {
	mutation := newVoteMutation(c.config, OpUpdateOne, withVote(v))
	return &VoteUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *VoteClient) UpdateOneID(id int) *VoteUpdateOne {
	mutation := newVoteMutation(c.config, OpUpdateOne, withVoteID(id))
	return &VoteUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Vote.
func (c *VoteClient) Delete() *VoteDelete {
	mutation := newVoteMutation(c.config, OpDelete)
	return &VoteDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *VoteClient) DeleteOne(v *Vote) *VoteDeleteOne {
	return c.DeleteOneID(v.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *VoteClient) DeleteOneID(id int) *VoteDeleteOne {
	builder := c.Delete().Where(vote.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &VoteDeleteOne{builder}
}

// Query returns a query builder for Vote.
func (c *VoteClient) Query() *VoteQuery {
	return &VoteQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeVote},
		inters: c.Interceptors(),
	}
}

// Get returns a Vote entity by its id.
func (c *VoteClient) Get(ctx context.Context, id int) (*Vote, error) {
	return c.Query().Where(vote.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *VoteClient) GetX(ctx context.Context, id int) *Vote {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *VoteClient) Hooks() []Hook {
	return c.hooks.Vote
}

// Interceptors returns the client interceptors.
func (c *VoteClient) Interceptors() []Interceptor {
	return c.inters.Vote
}

func (c *VoteClient) mutate(ctx context.Context, m *VoteMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&VoteCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&VoteUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&VoteUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&VoteDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Vote mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Author, Comment, Post, Vote []ent.Hook
	}
	inters struct {
		Author, Comment, Post, Vote []ent.Interceptor
	}
)
