// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package internal

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/believe-go"
	"github.com/stainless-sdks/believe-go/option"
	"github.com/stainless-sdks/believe-terraform/internal/services/biscuit"
	"github.com/stainless-sdks/believe-terraform/internal/services/character"
	"github.com/stainless-sdks/believe-terraform/internal/services/coaching_principle"
	"github.com/stainless-sdks/believe-terraform/internal/services/episode"
	"github.com/stainless-sdks/believe-terraform/internal/services/match"
	"github.com/stainless-sdks/believe-terraform/internal/services/pep_talk"
	"github.com/stainless-sdks/believe-terraform/internal/services/quote"
	"github.com/stainless-sdks/believe-terraform/internal/services/team"
	"github.com/stainless-sdks/believe-terraform/internal/services/team_member"
	"github.com/stainless-sdks/believe-terraform/internal/services/version"
	"github.com/stainless-sdks/believe-terraform/internal/services/webhook"
)

var _ provider.ProviderWithConfigValidators = (*BelieveProvider)(nil)

// BelieveProvider defines the provider implementation.
type BelieveProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// BelieveProviderModel describes the provider data model.
type BelieveProviderModel struct {
	BaseURL types.String `tfsdk:"base_url" json:"base_url,optional"`
	APIKey  types.String `tfsdk:"api_key" json:"api_key,optional"`
}

func (p *BelieveProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "believe"
	resp.Version = p.version
}

func ProviderSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"base_url": schema.StringAttribute{
				Description: "Set the base url that the provider connects to.",
				Optional:    true,
			},
			"api_key": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (p *BelieveProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = ProviderSchema(ctx)
}

func (p *BelieveProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	var data BelieveProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	opts := []option.RequestOption{}

	if !data.BaseURL.IsNull() && !data.BaseURL.IsUnknown() {
		opts = append(opts, option.WithBaseURL(data.BaseURL.ValueString()))
	} else if o, ok := os.LookupEnv("BELIEVE_BASE_URL"); ok {
		opts = append(opts, option.WithBaseURL(o))
	}

	if !data.APIKey.IsNull() && !data.APIKey.IsUnknown() {
		opts = append(opts, option.WithAPIKey(data.APIKey.ValueString()))
	} else if o, ok := os.LookupEnv("BELIEVE_API_KEY"); ok {
		opts = append(opts, option.WithAPIKey(o))
	} else {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"Missing api_key value",
			"The api_key field is required. Set it in provider configuration or via the \"BELIEVE_API_KEY\" environment variable.",
		)
		return
	}

	client := believe.NewClient(
		opts...,
	)

	resp.DataSourceData = &client
	resp.ResourceData = &client
}

func (p *BelieveProvider) ConfigValidators(_ context.Context) []provider.ConfigValidator {
	return []provider.ConfigValidator{}
}

func (p *BelieveProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		character.NewResource,
		team.NewResource,
		match.NewResource,
		episode.NewResource,
		quote.NewResource,
		team_member.NewResource,
		webhook.NewResource,
	}
}

func (p *BelieveProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		character.NewCharacterDataSource,
		character.NewCharactersDataSource,
		team.NewTeamDataSource,
		team.NewTeamsDataSource,
		match.NewMatchDataSource,
		match.NewMatchesDataSource,
		episode.NewEpisodeDataSource,
		episode.NewEpisodesDataSource,
		quote.NewQuoteDataSource,
		quote.NewQuotesDataSource,
		coaching_principle.NewCoachingPrincipleDataSource,
		coaching_principle.NewCoachingPrinciplesDataSource,
		biscuit.NewBiscuitDataSource,
		biscuit.NewBiscuitsDataSource,
		pep_talk.NewPepTalkDataSource,
		team_member.NewTeamMemberDataSource,
		team_member.NewTeamMembersDataSource,
		webhook.NewWebhookDataSource,
		version.NewVersionDataSource,
	}
}

func NewProvider(version string) func() provider.Provider {
	return func() provider.Provider {
		return &BelieveProvider{
			version: version,
		}
	}
}
