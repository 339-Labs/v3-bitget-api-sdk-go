/*
 * Bitget Open API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 2.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package com.bitget.openapi.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.JsonArray;
import com.google.gson.JsonDeserializationContext;
import com.google.gson.JsonDeserializer;
import com.google.gson.JsonElement;
import com.google.gson.JsonObject;
import com.google.gson.JsonParseException;
import com.google.gson.TypeAdapterFactory;
import com.google.gson.reflect.TypeToken;

import java.lang.reflect.Type;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Map.Entry;
import java.util.Set;

import com.bitget.openapi.JSON;

/**
 * MarginCrossVipResult
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class MarginCrossVipResult {
  public static final String SERIALIZED_NAME_DAILY_INTEREST_RATE = "dailyInterestRate";
  @SerializedName(SERIALIZED_NAME_DAILY_INTEREST_RATE)
  private String dailyInterestRate;

  public static final String SERIALIZED_NAME_DISCOUNT_RATE = "discountRate";
  @SerializedName(SERIALIZED_NAME_DISCOUNT_RATE)
  private String discountRate;

  public static final String SERIALIZED_NAME_LEVEL = "level";
  @SerializedName(SERIALIZED_NAME_LEVEL)
  private String level;

  public static final String SERIALIZED_NAME_YEARLY_INTEREST_RATE = "yearlyInterestRate";
  @SerializedName(SERIALIZED_NAME_YEARLY_INTEREST_RATE)
  private String yearlyInterestRate;

  public MarginCrossVipResult() {
  }

  public MarginCrossVipResult dailyInterestRate(String dailyInterestRate) {
    
    this.dailyInterestRate = dailyInterestRate;
    return this;
  }

   /**
   * Get dailyInterestRate
   * @return dailyInterestRate
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDailyInterestRate() {
    return dailyInterestRate;
  }


  public void setDailyInterestRate(String dailyInterestRate) {
    this.dailyInterestRate = dailyInterestRate;
  }


  public MarginCrossVipResult discountRate(String discountRate) {
    
    this.discountRate = discountRate;
    return this;
  }

   /**
   * Get discountRate
   * @return discountRate
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDiscountRate() {
    return discountRate;
  }


  public void setDiscountRate(String discountRate) {
    this.discountRate = discountRate;
  }


  public MarginCrossVipResult level(String level) {
    
    this.level = level;
    return this;
  }

   /**
   * Get level
   * @return level
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getLevel() {
    return level;
  }


  public void setLevel(String level) {
    this.level = level;
  }


  public MarginCrossVipResult yearlyInterestRate(String yearlyInterestRate) {
    
    this.yearlyInterestRate = yearlyInterestRate;
    return this;
  }

   /**
   * Get yearlyInterestRate
   * @return yearlyInterestRate
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getYearlyInterestRate() {
    return yearlyInterestRate;
  }


  public void setYearlyInterestRate(String yearlyInterestRate) {
    this.yearlyInterestRate = yearlyInterestRate;
  }

  /**
   * A container for additional, undeclared properties.
   * This is a holder for any undeclared properties as specified with
   * the 'additionalProperties' keyword in the OAS document.
   */
  private Map<String, Object> additionalProperties;

  /**
   * Set the additional (undeclared) property with the specified name and value.
   * If the property does not already exist, create it otherwise replace it.
   *
   * @param key name of the property
   * @param value value of the property
   * @return the MarginCrossVipResult instance itself
   */
  public MarginCrossVipResult putAdditionalProperty(String key, Object value) {
    if (this.additionalProperties == null) {
        this.additionalProperties = new HashMap<String, Object>();
    }
    this.additionalProperties.put(key, value);
    return this;
  }

  /**
   * Return the additional (undeclared) property.
   *
   * @return a map of objects
   */
  public Map<String, Object> getAdditionalProperties() {
    return additionalProperties;
  }

  /**
   * Return the additional (undeclared) property with the specified name.
   *
   * @param key name of the property
   * @return an object
   */
  public Object getAdditionalProperty(String key) {
    if (this.additionalProperties == null) {
        return null;
    }
    return this.additionalProperties.get(key);
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    MarginCrossVipResult marginCrossVipResult = (MarginCrossVipResult) o;
    return Objects.equals(this.dailyInterestRate, marginCrossVipResult.dailyInterestRate) &&
        Objects.equals(this.discountRate, marginCrossVipResult.discountRate) &&
        Objects.equals(this.level, marginCrossVipResult.level) &&
        Objects.equals(this.yearlyInterestRate, marginCrossVipResult.yearlyInterestRate)&&
        Objects.equals(this.additionalProperties, marginCrossVipResult.additionalProperties);
  }

  @Override
  public int hashCode() {
    return Objects.hash(dailyInterestRate, discountRate, level, yearlyInterestRate, additionalProperties);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class MarginCrossVipResult {\n");
    sb.append("    dailyInterestRate: ").append(toIndentedString(dailyInterestRate)).append("\n");
    sb.append("    discountRate: ").append(toIndentedString(discountRate)).append("\n");
    sb.append("    level: ").append(toIndentedString(level)).append("\n");
    sb.append("    yearlyInterestRate: ").append(toIndentedString(yearlyInterestRate)).append("\n");
    sb.append("    additionalProperties: ").append(toIndentedString(additionalProperties)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }


  public static HashSet<String> openapiFields;
  public static HashSet<String> openapiRequiredFields;

  static {
    // a set of all properties/fields (JSON key names)
    openapiFields = new HashSet<String>();
    openapiFields.add("dailyInterestRate");
    openapiFields.add("discountRate");
    openapiFields.add("level");
    openapiFields.add("yearlyInterestRate");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to MarginCrossVipResult
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!MarginCrossVipResult.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in MarginCrossVipResult is not found in the empty JSON string", MarginCrossVipResult.openapiRequiredFields.toString()));
        }
      }
      if ((jsonObj.get("dailyInterestRate") != null && !jsonObj.get("dailyInterestRate").isJsonNull()) && !jsonObj.get("dailyInterestRate").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `dailyInterestRate` to be a primitive type in the JSON string but got `%s`", jsonObj.get("dailyInterestRate").toString()));
      }
      if ((jsonObj.get("discountRate") != null && !jsonObj.get("discountRate").isJsonNull()) && !jsonObj.get("discountRate").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `discountRate` to be a primitive type in the JSON string but got `%s`", jsonObj.get("discountRate").toString()));
      }
      if ((jsonObj.get("level") != null && !jsonObj.get("level").isJsonNull()) && !jsonObj.get("level").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `level` to be a primitive type in the JSON string but got `%s`", jsonObj.get("level").toString()));
      }
      if ((jsonObj.get("yearlyInterestRate") != null && !jsonObj.get("yearlyInterestRate").isJsonNull()) && !jsonObj.get("yearlyInterestRate").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `yearlyInterestRate` to be a primitive type in the JSON string but got `%s`", jsonObj.get("yearlyInterestRate").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!MarginCrossVipResult.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'MarginCrossVipResult' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<MarginCrossVipResult> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(MarginCrossVipResult.class));

       return (TypeAdapter<T>) new TypeAdapter<MarginCrossVipResult>() {
           @Override
           public void write(JsonWriter out, MarginCrossVipResult value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             obj.remove("additionalProperties");
             // serialize additonal properties
             if (value.getAdditionalProperties() != null) {
               for (Map.Entry<String, Object> entry : value.getAdditionalProperties().entrySet()) {
                 if (entry.getValue() instanceof String)
                   obj.addProperty(entry.getKey(), (String) entry.getValue());
                 else if (entry.getValue() instanceof Number)
                   obj.addProperty(entry.getKey(), (Number) entry.getValue());
                 else if (entry.getValue() instanceof Boolean)
                   obj.addProperty(entry.getKey(), (Boolean) entry.getValue());
                 else if (entry.getValue() instanceof Character)
                   obj.addProperty(entry.getKey(), (Character) entry.getValue());
                 else {
                   obj.add(entry.getKey(), gson.toJsonTree(entry.getValue()).getAsJsonObject());
                 }
               }
             }
             elementAdapter.write(out, obj);
           }

           @Override
           public MarginCrossVipResult read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             // store additional fields in the deserialized instance
             MarginCrossVipResult instance = thisAdapter.fromJsonTree(jsonObj);
             for (Map.Entry<String, JsonElement> entry : jsonObj.entrySet()) {
               if (!openapiFields.contains(entry.getKey())) {
                 if (entry.getValue().isJsonPrimitive()) { // primitive type
                   if (entry.getValue().getAsJsonPrimitive().isString())
                     instance.putAdditionalProperty(entry.getKey(), entry.getValue().getAsString());
                   else if (entry.getValue().getAsJsonPrimitive().isNumber())
                     instance.putAdditionalProperty(entry.getKey(), entry.getValue().getAsNumber());
                   else if (entry.getValue().getAsJsonPrimitive().isBoolean())
                     instance.putAdditionalProperty(entry.getKey(), entry.getValue().getAsBoolean());
                   else
                     throw new IllegalArgumentException(String.format("The field `%s` has unknown primitive type. Value: %s", entry.getKey(), entry.getValue().toString()));
                 } else if (entry.getValue().isJsonArray()) {
                     instance.putAdditionalProperty(entry.getKey(), gson.fromJson(entry.getValue(), List.class));
                 } else { // JSON object
                     instance.putAdditionalProperty(entry.getKey(), gson.fromJson(entry.getValue(), HashMap.class));
                 }
               }
             }
             return instance;
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of MarginCrossVipResult given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of MarginCrossVipResult
  * @throws IOException if the JSON string is invalid with respect to MarginCrossVipResult
  */
  public static MarginCrossVipResult fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, MarginCrossVipResult.class);
  }

 /**
  * Convert an instance of MarginCrossVipResult to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

