package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/kubernetes/mcp-server/config"
	"github.com/kubernetes/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Readappsv1beta2namespacedstatefulsetHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["exact"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("exact=%v", val))
		}
		if val, ok := args["export"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("export=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/apis/apps/v1beta2/namespaces/%s/statefulsets/%s%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			req.Header.Set("authorization", cfg.APIKey)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.Iok8sapiappsv1beta2StatefulSet
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateReadappsv1beta2namespacedstatefulsetTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_apis_apps_v1beta2_namespaces_namespace_statefulsets_name",
		mcp.WithDescription("read the specified StatefulSet"),
		mcp.WithBoolean("exact", mcp.Description("Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'.")),
		mcp.WithBoolean("export", mcp.Description("Should this value be exported.  Export strips fields that a user can not specify.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Readappsv1beta2namespacedstatefulsetHandler(cfg),
	}
}
