# DefaultApi

All URIs are relative to *http://mgh.us-east-1.amazonaws.com*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**associateCreatedArtifact**](DefaultApi.md#associateCreatedArtifact) | **POST** /#X-Amz-Target&#x3D;AWSMigrationHub.AssociateCreatedArtifact |  |
| [**associateDiscoveredResource**](DefaultApi.md#associateDiscoveredResource) | **POST** /#X-Amz-Target&#x3D;AWSMigrationHub.AssociateDiscoveredResource |  |
| [**createProgressUpdateStream**](DefaultApi.md#createProgressUpdateStream) | **POST** /#X-Amz-Target&#x3D;AWSMigrationHub.CreateProgressUpdateStream |  |
| [**deleteProgressUpdateStream**](DefaultApi.md#deleteProgressUpdateStream) | **POST** /#X-Amz-Target&#x3D;AWSMigrationHub.DeleteProgressUpdateStream |  |
| [**describeApplicationState**](DefaultApi.md#describeApplicationState) | **POST** /#X-Amz-Target&#x3D;AWSMigrationHub.DescribeApplicationState |  |
| [**describeMigrationTask**](DefaultApi.md#describeMigrationTask) | **POST** /#X-Amz-Target&#x3D;AWSMigrationHub.DescribeMigrationTask |  |
| [**disassociateCreatedArtifact**](DefaultApi.md#disassociateCreatedArtifact) | **POST** /#X-Amz-Target&#x3D;AWSMigrationHub.DisassociateCreatedArtifact |  |
| [**disassociateDiscoveredResource**](DefaultApi.md#disassociateDiscoveredResource) | **POST** /#X-Amz-Target&#x3D;AWSMigrationHub.DisassociateDiscoveredResource |  |
| [**importMigrationTask**](DefaultApi.md#importMigrationTask) | **POST** /#X-Amz-Target&#x3D;AWSMigrationHub.ImportMigrationTask |  |
| [**listApplicationStates**](DefaultApi.md#listApplicationStates) | **POST** /#X-Amz-Target&#x3D;AWSMigrationHub.ListApplicationStates |  |
| [**listCreatedArtifacts**](DefaultApi.md#listCreatedArtifacts) | **POST** /#X-Amz-Target&#x3D;AWSMigrationHub.ListCreatedArtifacts |  |
| [**listDiscoveredResources**](DefaultApi.md#listDiscoveredResources) | **POST** /#X-Amz-Target&#x3D;AWSMigrationHub.ListDiscoveredResources |  |
| [**listMigrationTasks**](DefaultApi.md#listMigrationTasks) | **POST** /#X-Amz-Target&#x3D;AWSMigrationHub.ListMigrationTasks |  |
| [**listProgressUpdateStreams**](DefaultApi.md#listProgressUpdateStreams) | **POST** /#X-Amz-Target&#x3D;AWSMigrationHub.ListProgressUpdateStreams |  |
| [**notifyApplicationState**](DefaultApi.md#notifyApplicationState) | **POST** /#X-Amz-Target&#x3D;AWSMigrationHub.NotifyApplicationState |  |
| [**notifyMigrationTaskState**](DefaultApi.md#notifyMigrationTaskState) | **POST** /#X-Amz-Target&#x3D;AWSMigrationHub.NotifyMigrationTaskState |  |
| [**putResourceAttributes**](DefaultApi.md#putResourceAttributes) | **POST** /#X-Amz-Target&#x3D;AWSMigrationHub.PutResourceAttributes |  |


<a id="associateCreatedArtifact"></a>
# **associateCreatedArtifact**
> Object associateCreatedArtifact(xAmzTarget, associateCreatedArtifactRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



&lt;p&gt;Associates a created artifact of an AWS cloud resource, the target receiving the migration, with the migration task performed by a migration tool. This API has the following traits:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt;Migration tools can call the &lt;code&gt;AssociateCreatedArtifact&lt;/code&gt; operation to indicate which AWS artifact is associated with a migration task.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;The created artifact name must be provided in ARN (Amazon Resource Name) format which will contain information about type and region; for example: &lt;code&gt;arn:aws:ec2:us-east-1:488216288981:image/ami-6d0ba87b&lt;/code&gt;.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;Examples of the AWS resource behind the created artifact are, AMI&#39;s, EC2 instance, or DMS endpoint, etc.&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt;

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://mgh.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String xAmzTarget = "AWSMigrationHub.AssociateCreatedArtifact"; // String | 
    AssociateCreatedArtifactRequest associateCreatedArtifactRequest = new AssociateCreatedArtifactRequest(); // AssociateCreatedArtifactRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      Object result = apiInstance.associateCreatedArtifact(xAmzTarget, associateCreatedArtifactRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#associateCreatedArtifact");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **xAmzTarget** | **String**|  | [enum: AWSMigrationHub.AssociateCreatedArtifact] |
| **associateCreatedArtifactRequest** | [**AssociateCreatedArtifactRequest**](AssociateCreatedArtifactRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

**Object**

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | AccessDeniedException |  -  |
| **481** | ThrottlingException |  -  |
| **482** | InternalServerError |  -  |
| **483** | ServiceUnavailableException |  -  |
| **484** | DryRunOperation |  -  |
| **485** | UnauthorizedOperation |  -  |
| **486** | InvalidInputException |  -  |
| **487** | ResourceNotFoundException |  -  |
| **488** | HomeRegionNotSetException |  -  |

<a id="associateDiscoveredResource"></a>
# **associateDiscoveredResource**
> Object associateDiscoveredResource(xAmzTarget, associateDiscoveredResourceRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Associates a discovered resource ID from Application Discovery Service with a migration task.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://mgh.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String xAmzTarget = "AWSMigrationHub.AssociateDiscoveredResource"; // String | 
    AssociateDiscoveredResourceRequest associateDiscoveredResourceRequest = new AssociateDiscoveredResourceRequest(); // AssociateDiscoveredResourceRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      Object result = apiInstance.associateDiscoveredResource(xAmzTarget, associateDiscoveredResourceRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#associateDiscoveredResource");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **xAmzTarget** | **String**|  | [enum: AWSMigrationHub.AssociateDiscoveredResource] |
| **associateDiscoveredResourceRequest** | [**AssociateDiscoveredResourceRequest**](AssociateDiscoveredResourceRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

**Object**

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | AccessDeniedException |  -  |
| **481** | ThrottlingException |  -  |
| **482** | InternalServerError |  -  |
| **483** | ServiceUnavailableException |  -  |
| **484** | DryRunOperation |  -  |
| **485** | UnauthorizedOperation |  -  |
| **486** | InvalidInputException |  -  |
| **487** | PolicyErrorException |  -  |
| **488** | ResourceNotFoundException |  -  |
| **489** | HomeRegionNotSetException |  -  |

<a id="createProgressUpdateStream"></a>
# **createProgressUpdateStream**
> Object createProgressUpdateStream(xAmzTarget, createProgressUpdateStreamRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Creates a progress update stream which is an AWS resource used for access control as well as a namespace for migration task names that is implicitly linked to your AWS account. It must uniquely identify the migration tool as it is used for all updates made by the tool; however, it does not need to be unique for each AWS account because it is scoped to the AWS account.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://mgh.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String xAmzTarget = "AWSMigrationHub.CreateProgressUpdateStream"; // String | 
    CreateProgressUpdateStreamRequest createProgressUpdateStreamRequest = new CreateProgressUpdateStreamRequest(); // CreateProgressUpdateStreamRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      Object result = apiInstance.createProgressUpdateStream(xAmzTarget, createProgressUpdateStreamRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#createProgressUpdateStream");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **xAmzTarget** | **String**|  | [enum: AWSMigrationHub.CreateProgressUpdateStream] |
| **createProgressUpdateStreamRequest** | [**CreateProgressUpdateStreamRequest**](CreateProgressUpdateStreamRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

**Object**

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | AccessDeniedException |  -  |
| **481** | ThrottlingException |  -  |
| **482** | InternalServerError |  -  |
| **483** | ServiceUnavailableException |  -  |
| **484** | DryRunOperation |  -  |
| **485** | UnauthorizedOperation |  -  |
| **486** | InvalidInputException |  -  |
| **487** | HomeRegionNotSetException |  -  |

<a id="deleteProgressUpdateStream"></a>
# **deleteProgressUpdateStream**
> Object deleteProgressUpdateStream(xAmzTarget, deleteProgressUpdateStreamRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



&lt;p&gt;Deletes a progress update stream, including all of its tasks, which was previously created as an AWS resource used for access control. This API has the following traits:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt;The only parameter needed for &lt;code&gt;DeleteProgressUpdateStream&lt;/code&gt; is the stream name (same as a &lt;code&gt;CreateProgressUpdateStream&lt;/code&gt; call).&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;The call will return, and a background process will asynchronously delete the stream and all of its resources (tasks, associated resources, resource attributes, created artifacts).&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;If the stream takes time to be deleted, it might still show up on a &lt;code&gt;ListProgressUpdateStreams&lt;/code&gt; call.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;CreateProgressUpdateStream&lt;/code&gt;, &lt;code&gt;ImportMigrationTask&lt;/code&gt;, &lt;code&gt;NotifyMigrationTaskState&lt;/code&gt;, and all Associate[*] APIs related to the tasks belonging to the stream will throw \&quot;InvalidInputException\&quot; if the stream of the same name is in the process of being deleted.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;Once the stream and all of its resources are deleted, &lt;code&gt;CreateProgressUpdateStream&lt;/code&gt; for a stream of the same name will succeed, and that stream will be an entirely new logical resource (without any resources associated with the old stream).&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt;

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://mgh.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String xAmzTarget = "AWSMigrationHub.DeleteProgressUpdateStream"; // String | 
    DeleteProgressUpdateStreamRequest deleteProgressUpdateStreamRequest = new DeleteProgressUpdateStreamRequest(); // DeleteProgressUpdateStreamRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      Object result = apiInstance.deleteProgressUpdateStream(xAmzTarget, deleteProgressUpdateStreamRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#deleteProgressUpdateStream");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **xAmzTarget** | **String**|  | [enum: AWSMigrationHub.DeleteProgressUpdateStream] |
| **deleteProgressUpdateStreamRequest** | [**DeleteProgressUpdateStreamRequest**](DeleteProgressUpdateStreamRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

**Object**

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | AccessDeniedException |  -  |
| **481** | ThrottlingException |  -  |
| **482** | InternalServerError |  -  |
| **483** | ServiceUnavailableException |  -  |
| **484** | DryRunOperation |  -  |
| **485** | UnauthorizedOperation |  -  |
| **486** | InvalidInputException |  -  |
| **487** | ResourceNotFoundException |  -  |
| **488** | HomeRegionNotSetException |  -  |

<a id="describeApplicationState"></a>
# **describeApplicationState**
> DescribeApplicationStateResult describeApplicationState(xAmzTarget, describeApplicationStateRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Gets the migration status of an application.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://mgh.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String xAmzTarget = "AWSMigrationHub.DescribeApplicationState"; // String | 
    DescribeApplicationStateRequest describeApplicationStateRequest = new DescribeApplicationStateRequest(); // DescribeApplicationStateRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      DescribeApplicationStateResult result = apiInstance.describeApplicationState(xAmzTarget, describeApplicationStateRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#describeApplicationState");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **xAmzTarget** | **String**|  | [enum: AWSMigrationHub.DescribeApplicationState] |
| **describeApplicationStateRequest** | [**DescribeApplicationStateRequest**](DescribeApplicationStateRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

[**DescribeApplicationStateResult**](DescribeApplicationStateResult.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | AccessDeniedException |  -  |
| **481** | ThrottlingException |  -  |
| **482** | InternalServerError |  -  |
| **483** | ServiceUnavailableException |  -  |
| **484** | InvalidInputException |  -  |
| **485** | PolicyErrorException |  -  |
| **486** | ResourceNotFoundException |  -  |
| **487** | HomeRegionNotSetException |  -  |

<a id="describeMigrationTask"></a>
# **describeMigrationTask**
> DescribeMigrationTaskResult describeMigrationTask(xAmzTarget, describeMigrationTaskRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Retrieves a list of all attributes associated with a specific migration task.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://mgh.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String xAmzTarget = "AWSMigrationHub.DescribeMigrationTask"; // String | 
    DescribeMigrationTaskRequest describeMigrationTaskRequest = new DescribeMigrationTaskRequest(); // DescribeMigrationTaskRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      DescribeMigrationTaskResult result = apiInstance.describeMigrationTask(xAmzTarget, describeMigrationTaskRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#describeMigrationTask");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **xAmzTarget** | **String**|  | [enum: AWSMigrationHub.DescribeMigrationTask] |
| **describeMigrationTaskRequest** | [**DescribeMigrationTaskRequest**](DescribeMigrationTaskRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

[**DescribeMigrationTaskResult**](DescribeMigrationTaskResult.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | AccessDeniedException |  -  |
| **481** | ThrottlingException |  -  |
| **482** | InternalServerError |  -  |
| **483** | ServiceUnavailableException |  -  |
| **484** | InvalidInputException |  -  |
| **485** | ResourceNotFoundException |  -  |
| **486** | HomeRegionNotSetException |  -  |

<a id="disassociateCreatedArtifact"></a>
# **disassociateCreatedArtifact**
> Object disassociateCreatedArtifact(xAmzTarget, disassociateCreatedArtifactRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



&lt;p&gt;Disassociates a created artifact of an AWS resource with a migration task performed by a migration tool that was previously associated. This API has the following traits:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt;A migration user can call the &lt;code&gt;DisassociateCreatedArtifacts&lt;/code&gt; operation to disassociate a created AWS Artifact from a migration task.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;The created artifact name must be provided in ARN (Amazon Resource Name) format which will contain information about type and region; for example: &lt;code&gt;arn:aws:ec2:us-east-1:488216288981:image/ami-6d0ba87b&lt;/code&gt;.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;Examples of the AWS resource behind the created artifact are, AMI&#39;s, EC2 instance, or RDS instance, etc.&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt;

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://mgh.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String xAmzTarget = "AWSMigrationHub.DisassociateCreatedArtifact"; // String | 
    DisassociateCreatedArtifactRequest disassociateCreatedArtifactRequest = new DisassociateCreatedArtifactRequest(); // DisassociateCreatedArtifactRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      Object result = apiInstance.disassociateCreatedArtifact(xAmzTarget, disassociateCreatedArtifactRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#disassociateCreatedArtifact");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **xAmzTarget** | **String**|  | [enum: AWSMigrationHub.DisassociateCreatedArtifact] |
| **disassociateCreatedArtifactRequest** | [**DisassociateCreatedArtifactRequest**](DisassociateCreatedArtifactRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

**Object**

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | AccessDeniedException |  -  |
| **481** | ThrottlingException |  -  |
| **482** | InternalServerError |  -  |
| **483** | ServiceUnavailableException |  -  |
| **484** | DryRunOperation |  -  |
| **485** | UnauthorizedOperation |  -  |
| **486** | InvalidInputException |  -  |
| **487** | ResourceNotFoundException |  -  |
| **488** | HomeRegionNotSetException |  -  |

<a id="disassociateDiscoveredResource"></a>
# **disassociateDiscoveredResource**
> Object disassociateDiscoveredResource(xAmzTarget, disassociateDiscoveredResourceRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Disassociate an Application Discovery Service discovered resource from a migration task.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://mgh.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String xAmzTarget = "AWSMigrationHub.DisassociateDiscoveredResource"; // String | 
    DisassociateDiscoveredResourceRequest disassociateDiscoveredResourceRequest = new DisassociateDiscoveredResourceRequest(); // DisassociateDiscoveredResourceRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      Object result = apiInstance.disassociateDiscoveredResource(xAmzTarget, disassociateDiscoveredResourceRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#disassociateDiscoveredResource");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **xAmzTarget** | **String**|  | [enum: AWSMigrationHub.DisassociateDiscoveredResource] |
| **disassociateDiscoveredResourceRequest** | [**DisassociateDiscoveredResourceRequest**](DisassociateDiscoveredResourceRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

**Object**

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | AccessDeniedException |  -  |
| **481** | ThrottlingException |  -  |
| **482** | InternalServerError |  -  |
| **483** | ServiceUnavailableException |  -  |
| **484** | DryRunOperation |  -  |
| **485** | UnauthorizedOperation |  -  |
| **486** | InvalidInputException |  -  |
| **487** | ResourceNotFoundException |  -  |
| **488** | HomeRegionNotSetException |  -  |

<a id="importMigrationTask"></a>
# **importMigrationTask**
> Object importMigrationTask(xAmzTarget, importMigrationTaskRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



&lt;p&gt;Registers a new migration task which represents a server, database, etc., being migrated to AWS by a migration tool.&lt;/p&gt; &lt;p&gt;This API is a prerequisite to calling the &lt;code&gt;NotifyMigrationTaskState&lt;/code&gt; API as the migration tool must first register the migration task with Migration Hub.&lt;/p&gt;

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://mgh.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String xAmzTarget = "AWSMigrationHub.ImportMigrationTask"; // String | 
    ImportMigrationTaskRequest importMigrationTaskRequest = new ImportMigrationTaskRequest(); // ImportMigrationTaskRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      Object result = apiInstance.importMigrationTask(xAmzTarget, importMigrationTaskRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#importMigrationTask");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **xAmzTarget** | **String**|  | [enum: AWSMigrationHub.ImportMigrationTask] |
| **importMigrationTaskRequest** | [**ImportMigrationTaskRequest**](ImportMigrationTaskRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

**Object**

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | AccessDeniedException |  -  |
| **481** | ThrottlingException |  -  |
| **482** | InternalServerError |  -  |
| **483** | ServiceUnavailableException |  -  |
| **484** | DryRunOperation |  -  |
| **485** | UnauthorizedOperation |  -  |
| **486** | InvalidInputException |  -  |
| **487** | ResourceNotFoundException |  -  |
| **488** | HomeRegionNotSetException |  -  |

<a id="listApplicationStates"></a>
# **listApplicationStates**
> ListApplicationStatesResult listApplicationStates(xAmzTarget, listApplicationStatesRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, maxResults, nextToken)



Lists all the migration statuses for your applications. If you use the optional &lt;code&gt;ApplicationIds&lt;/code&gt; parameter, only the migration statuses for those applications will be returned.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://mgh.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String xAmzTarget = "AWSMigrationHub.ListApplicationStates"; // String | 
    ListApplicationStatesRequest listApplicationStatesRequest = new ListApplicationStatesRequest(); // ListApplicationStatesRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    String maxResults = "maxResults_example"; // String | Pagination limit
    String nextToken = "nextToken_example"; // String | Pagination token
    try {
      ListApplicationStatesResult result = apiInstance.listApplicationStates(xAmzTarget, listApplicationStatesRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, maxResults, nextToken);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#listApplicationStates");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **xAmzTarget** | **String**|  | [enum: AWSMigrationHub.ListApplicationStates] |
| **listApplicationStatesRequest** | [**ListApplicationStatesRequest**](ListApplicationStatesRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |
| **maxResults** | **String**| Pagination limit | [optional] |
| **nextToken** | **String**| Pagination token | [optional] |

### Return type

[**ListApplicationStatesResult**](ListApplicationStatesResult.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | AccessDeniedException |  -  |
| **481** | ThrottlingException |  -  |
| **482** | InternalServerError |  -  |
| **483** | ServiceUnavailableException |  -  |
| **484** | InvalidInputException |  -  |
| **485** | HomeRegionNotSetException |  -  |

<a id="listCreatedArtifacts"></a>
# **listCreatedArtifacts**
> ListCreatedArtifactsResult listCreatedArtifacts(xAmzTarget, listCreatedArtifactsRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, maxResults, nextToken)



&lt;p&gt;Lists the created artifacts attached to a given migration task in an update stream. This API has the following traits:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt;Gets the list of the created artifacts while migration is taking place.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;Shows the artifacts created by the migration tool that was associated by the &lt;code&gt;AssociateCreatedArtifact&lt;/code&gt; API. &lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;Lists created artifacts in a paginated interface. &lt;/p&gt; &lt;/li&gt; &lt;/ul&gt;

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://mgh.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String xAmzTarget = "AWSMigrationHub.ListCreatedArtifacts"; // String | 
    ListCreatedArtifactsRequest listCreatedArtifactsRequest = new ListCreatedArtifactsRequest(); // ListCreatedArtifactsRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    String maxResults = "maxResults_example"; // String | Pagination limit
    String nextToken = "nextToken_example"; // String | Pagination token
    try {
      ListCreatedArtifactsResult result = apiInstance.listCreatedArtifacts(xAmzTarget, listCreatedArtifactsRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, maxResults, nextToken);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#listCreatedArtifacts");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **xAmzTarget** | **String**|  | [enum: AWSMigrationHub.ListCreatedArtifacts] |
| **listCreatedArtifactsRequest** | [**ListCreatedArtifactsRequest**](ListCreatedArtifactsRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |
| **maxResults** | **String**| Pagination limit | [optional] |
| **nextToken** | **String**| Pagination token | [optional] |

### Return type

[**ListCreatedArtifactsResult**](ListCreatedArtifactsResult.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | AccessDeniedException |  -  |
| **481** | ThrottlingException |  -  |
| **482** | InternalServerError |  -  |
| **483** | ServiceUnavailableException |  -  |
| **484** | InvalidInputException |  -  |
| **485** | ResourceNotFoundException |  -  |
| **486** | HomeRegionNotSetException |  -  |

<a id="listDiscoveredResources"></a>
# **listDiscoveredResources**
> ListDiscoveredResourcesResult listDiscoveredResources(xAmzTarget, listDiscoveredResourcesRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, maxResults, nextToken)



Lists discovered resources associated with the given &lt;code&gt;MigrationTask&lt;/code&gt;.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://mgh.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String xAmzTarget = "AWSMigrationHub.ListDiscoveredResources"; // String | 
    ListDiscoveredResourcesRequest listDiscoveredResourcesRequest = new ListDiscoveredResourcesRequest(); // ListDiscoveredResourcesRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    String maxResults = "maxResults_example"; // String | Pagination limit
    String nextToken = "nextToken_example"; // String | Pagination token
    try {
      ListDiscoveredResourcesResult result = apiInstance.listDiscoveredResources(xAmzTarget, listDiscoveredResourcesRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, maxResults, nextToken);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#listDiscoveredResources");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **xAmzTarget** | **String**|  | [enum: AWSMigrationHub.ListDiscoveredResources] |
| **listDiscoveredResourcesRequest** | [**ListDiscoveredResourcesRequest**](ListDiscoveredResourcesRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |
| **maxResults** | **String**| Pagination limit | [optional] |
| **nextToken** | **String**| Pagination token | [optional] |

### Return type

[**ListDiscoveredResourcesResult**](ListDiscoveredResourcesResult.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | AccessDeniedException |  -  |
| **481** | ThrottlingException |  -  |
| **482** | InternalServerError |  -  |
| **483** | ServiceUnavailableException |  -  |
| **484** | InvalidInputException |  -  |
| **485** | ResourceNotFoundException |  -  |
| **486** | HomeRegionNotSetException |  -  |

<a id="listMigrationTasks"></a>
# **listMigrationTasks**
> ListMigrationTasksResult listMigrationTasks(xAmzTarget, listMigrationTasksRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, maxResults, nextToken)



&lt;p&gt;Lists all, or filtered by resource name, migration tasks associated with the user account making this call. This API has the following traits:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt;Can show a summary list of the most recent migration tasks.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;Can show a summary list of migration tasks associated with a given discovered resource.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;Lists migration tasks in a paginated interface.&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt;

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://mgh.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String xAmzTarget = "AWSMigrationHub.ListMigrationTasks"; // String | 
    ListMigrationTasksRequest listMigrationTasksRequest = new ListMigrationTasksRequest(); // ListMigrationTasksRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    String maxResults = "maxResults_example"; // String | Pagination limit
    String nextToken = "nextToken_example"; // String | Pagination token
    try {
      ListMigrationTasksResult result = apiInstance.listMigrationTasks(xAmzTarget, listMigrationTasksRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, maxResults, nextToken);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#listMigrationTasks");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **xAmzTarget** | **String**|  | [enum: AWSMigrationHub.ListMigrationTasks] |
| **listMigrationTasksRequest** | [**ListMigrationTasksRequest**](ListMigrationTasksRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |
| **maxResults** | **String**| Pagination limit | [optional] |
| **nextToken** | **String**| Pagination token | [optional] |

### Return type

[**ListMigrationTasksResult**](ListMigrationTasksResult.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | AccessDeniedException |  -  |
| **481** | ThrottlingException |  -  |
| **482** | InternalServerError |  -  |
| **483** | ServiceUnavailableException |  -  |
| **484** | InvalidInputException |  -  |
| **485** | PolicyErrorException |  -  |
| **486** | ResourceNotFoundException |  -  |
| **487** | HomeRegionNotSetException |  -  |

<a id="listProgressUpdateStreams"></a>
# **listProgressUpdateStreams**
> ListProgressUpdateStreamsResult listProgressUpdateStreams(xAmzTarget, listProgressUpdateStreamsRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, maxResults, nextToken)



Lists progress update streams associated with the user account making this call.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://mgh.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String xAmzTarget = "AWSMigrationHub.ListProgressUpdateStreams"; // String | 
    ListProgressUpdateStreamsRequest listProgressUpdateStreamsRequest = new ListProgressUpdateStreamsRequest(); // ListProgressUpdateStreamsRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    String maxResults = "maxResults_example"; // String | Pagination limit
    String nextToken = "nextToken_example"; // String | Pagination token
    try {
      ListProgressUpdateStreamsResult result = apiInstance.listProgressUpdateStreams(xAmzTarget, listProgressUpdateStreamsRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders, maxResults, nextToken);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#listProgressUpdateStreams");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **xAmzTarget** | **String**|  | [enum: AWSMigrationHub.ListProgressUpdateStreams] |
| **listProgressUpdateStreamsRequest** | [**ListProgressUpdateStreamsRequest**](ListProgressUpdateStreamsRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |
| **maxResults** | **String**| Pagination limit | [optional] |
| **nextToken** | **String**| Pagination token | [optional] |

### Return type

[**ListProgressUpdateStreamsResult**](ListProgressUpdateStreamsResult.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | AccessDeniedException |  -  |
| **481** | ThrottlingException |  -  |
| **482** | InternalServerError |  -  |
| **483** | ServiceUnavailableException |  -  |
| **484** | InvalidInputException |  -  |
| **485** | HomeRegionNotSetException |  -  |

<a id="notifyApplicationState"></a>
# **notifyApplicationState**
> Object notifyApplicationState(xAmzTarget, notifyApplicationStateRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



Sets the migration state of an application. For a given application identified by the value passed to &lt;code&gt;ApplicationId&lt;/code&gt;, its status is set or updated by passing one of three values to &lt;code&gt;Status&lt;/code&gt;: &lt;code&gt;NOT_STARTED | IN_PROGRESS | COMPLETED&lt;/code&gt;.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://mgh.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String xAmzTarget = "AWSMigrationHub.NotifyApplicationState"; // String | 
    NotifyApplicationStateRequest notifyApplicationStateRequest = new NotifyApplicationStateRequest(); // NotifyApplicationStateRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      Object result = apiInstance.notifyApplicationState(xAmzTarget, notifyApplicationStateRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#notifyApplicationState");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **xAmzTarget** | **String**|  | [enum: AWSMigrationHub.NotifyApplicationState] |
| **notifyApplicationStateRequest** | [**NotifyApplicationStateRequest**](NotifyApplicationStateRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

**Object**

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | AccessDeniedException |  -  |
| **481** | ThrottlingException |  -  |
| **482** | InternalServerError |  -  |
| **483** | ServiceUnavailableException |  -  |
| **484** | DryRunOperation |  -  |
| **485** | UnauthorizedOperation |  -  |
| **486** | InvalidInputException |  -  |
| **487** | PolicyErrorException |  -  |
| **488** | ResourceNotFoundException |  -  |
| **489** | HomeRegionNotSetException |  -  |

<a id="notifyMigrationTaskState"></a>
# **notifyMigrationTaskState**
> Object notifyMigrationTaskState(xAmzTarget, notifyMigrationTaskStateRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



&lt;p&gt;Notifies Migration Hub of the current status, progress, or other detail regarding a migration task. This API has the following traits:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt;Migration tools will call the &lt;code&gt;NotifyMigrationTaskState&lt;/code&gt; API to share the latest progress and status.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;MigrationTaskName&lt;/code&gt; is used for addressing updates to the correct target.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;ProgressUpdateStream&lt;/code&gt; is used for access control and to provide a namespace for each migration tool.&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt;

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://mgh.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String xAmzTarget = "AWSMigrationHub.NotifyMigrationTaskState"; // String | 
    NotifyMigrationTaskStateRequest notifyMigrationTaskStateRequest = new NotifyMigrationTaskStateRequest(); // NotifyMigrationTaskStateRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      Object result = apiInstance.notifyMigrationTaskState(xAmzTarget, notifyMigrationTaskStateRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#notifyMigrationTaskState");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **xAmzTarget** | **String**|  | [enum: AWSMigrationHub.NotifyMigrationTaskState] |
| **notifyMigrationTaskStateRequest** | [**NotifyMigrationTaskStateRequest**](NotifyMigrationTaskStateRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

**Object**

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | AccessDeniedException |  -  |
| **481** | ThrottlingException |  -  |
| **482** | InternalServerError |  -  |
| **483** | ServiceUnavailableException |  -  |
| **484** | DryRunOperation |  -  |
| **485** | UnauthorizedOperation |  -  |
| **486** | InvalidInputException |  -  |
| **487** | ResourceNotFoundException |  -  |
| **488** | HomeRegionNotSetException |  -  |

<a id="putResourceAttributes"></a>
# **putResourceAttributes**
> Object putResourceAttributes(xAmzTarget, putResourceAttributesRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders)



&lt;p&gt;Provides identifying details of the resource being migrated so that it can be associated in the Application Discovery Service repository. This association occurs asynchronously after &lt;code&gt;PutResourceAttributes&lt;/code&gt; returns.&lt;/p&gt; &lt;important&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt;Keep in mind that subsequent calls to PutResourceAttributes will override previously stored attributes. For example, if it is first called with a MAC address, but later, it is desired to &lt;i&gt;add&lt;/i&gt; an IP address, it will then be required to call it with &lt;i&gt;both&lt;/i&gt; the IP and MAC addresses to prevent overriding the MAC address.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;Note the instructions regarding the special use case of the &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/migrationhub/latest/ug/API_PutResourceAttributes.html#migrationhub-PutResourceAttributes-request-ResourceAttributeList\&quot;&gt; &lt;code&gt;ResourceAttributeList&lt;/code&gt; &lt;/a&gt; parameter when specifying any \&quot;VM\&quot; related value.&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; &lt;/important&gt; &lt;note&gt; &lt;p&gt;Because this is an asynchronous call, it will always return 200, whether an association occurs or not. To confirm if an association was found based on the provided details, call &lt;code&gt;ListDiscoveredResources&lt;/code&gt;.&lt;/p&gt; &lt;/note&gt;

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://mgh.us-east-1.amazonaws.com");
    
    // Configure API key authorization: hmac
    ApiKeyAuth hmac = (ApiKeyAuth) defaultClient.getAuthentication("hmac");
    hmac.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //hmac.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String xAmzTarget = "AWSMigrationHub.PutResourceAttributes"; // String | 
    PutResourceAttributesRequest putResourceAttributesRequest = new PutResourceAttributesRequest(); // PutResourceAttributesRequest | 
    String xAmzContentSha256 = "xAmzContentSha256_example"; // String | 
    String xAmzDate = "xAmzDate_example"; // String | 
    String xAmzAlgorithm = "xAmzAlgorithm_example"; // String | 
    String xAmzCredential = "xAmzCredential_example"; // String | 
    String xAmzSecurityToken = "xAmzSecurityToken_example"; // String | 
    String xAmzSignature = "xAmzSignature_example"; // String | 
    String xAmzSignedHeaders = "xAmzSignedHeaders_example"; // String | 
    try {
      Object result = apiInstance.putResourceAttributes(xAmzTarget, putResourceAttributesRequest, xAmzContentSha256, xAmzDate, xAmzAlgorithm, xAmzCredential, xAmzSecurityToken, xAmzSignature, xAmzSignedHeaders);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#putResourceAttributes");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **xAmzTarget** | **String**|  | [enum: AWSMigrationHub.PutResourceAttributes] |
| **putResourceAttributesRequest** | [**PutResourceAttributesRequest**](PutResourceAttributesRequest.md)|  | |
| **xAmzContentSha256** | **String**|  | [optional] |
| **xAmzDate** | **String**|  | [optional] |
| **xAmzAlgorithm** | **String**|  | [optional] |
| **xAmzCredential** | **String**|  | [optional] |
| **xAmzSecurityToken** | **String**|  | [optional] |
| **xAmzSignature** | **String**|  | [optional] |
| **xAmzSignedHeaders** | **String**|  | [optional] |

### Return type

**Object**

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **480** | AccessDeniedException |  -  |
| **481** | ThrottlingException |  -  |
| **482** | InternalServerError |  -  |
| **483** | ServiceUnavailableException |  -  |
| **484** | DryRunOperation |  -  |
| **485** | UnauthorizedOperation |  -  |
| **486** | InvalidInputException |  -  |
| **487** | ResourceNotFoundException |  -  |
| **488** | HomeRegionNotSetException |  -  |

