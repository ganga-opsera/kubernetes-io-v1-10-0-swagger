package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/kubernetes/mcp-server/config"
	"github.com/kubernetes/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Createeventsv1beta1namespacedeventHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.Iok8sapieventsv1beta1Event
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/apis/events.k8s.io/v1beta1/namespaces/%s/events", cfg.BaseURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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
		var result models.Iok8sapieventsv1beta1Event
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

func CreateCreateeventsv1beta1namespacedeventTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_apis_events_k8s_io_v1beta1_namespaces_namespace_events",
		mcp.WithDescription("create an Event"),
		mcp.WithString("action", mcp.Description("Input parameter: What action was taken/failed regarding to the regarding object.")),
		mcp.WithString("kind", mcp.Description("Input parameter: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds")),
		mcp.WithString("deprecatedFirstTimestamp", mcp.Description("")),
		mcp.WithString("deprecatedLastTimestamp", mcp.Description("")),
		mcp.WithObject("metadata", mcp.Description("Input parameter: ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.")),
		mcp.WithString("note", mcp.Description("Input parameter: Optional. A human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.")),
		mcp.WithString("reason", mcp.Description("Input parameter: Why the action was taken.")),
		mcp.WithObject("regarding", mcp.Description("Input parameter: ObjectReference contains enough information to let you inspect or modify the referred object.")),
		mcp.WithString("reportingController", mcp.Description("Input parameter: Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.")),
		mcp.WithString("reportingInstance", mcp.Description("Input parameter: ID of the controller instance, e.g. `kubelet-xyzf`.")),
		mcp.WithObject("series", mcp.Description("Input parameter: EventSeries contain information on series of events, i.e. thing that was/is happening continuously for some time.")),
		mcp.WithString("type", mcp.Description("Input parameter: Type of this event (Normal, Warning), new types could be added in the future.")),
		mcp.WithObject("related", mcp.Description("Input parameter: ObjectReference contains enough information to let you inspect or modify the referred object.")),
		mcp.WithString("apiVersion", mcp.Description("Input parameter: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources")),
		mcp.WithString("eventTime", mcp.Required(), mcp.Description("")),
		mcp.WithNumber("deprecatedCount", mcp.Description("Input parameter: Deprecated field assuring backward compatibility with core.v1 Event type")),
		mcp.WithObject("deprecatedSource", mcp.Description("Input parameter: EventSource contains information for an event.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Createeventsv1beta1namespacedeventHandler(cfg),
	}
}
