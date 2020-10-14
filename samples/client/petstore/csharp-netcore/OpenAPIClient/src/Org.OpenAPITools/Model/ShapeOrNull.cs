/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * The version of the OpenAPI document: 1.0.0
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */


using System;
using System.Collections;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.IO;
using System.Runtime.Serialization;
using System.Text;
using System.Text.RegularExpressions;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;
using Newtonsoft.Json.Linq;
using JsonSubTypes;
using System.ComponentModel.DataAnnotations;
using OpenAPIDateConverter = Org.OpenAPITools.Client.OpenAPIDateConverter;
using OpenAPIClientUtils = Org.OpenAPITools.Client.ClientUtils;

namespace Org.OpenAPITools.Model
{
    /// <summary>
    /// The value may be a shape or the &#39;null&#39; value. This is introduced in OAS schema &gt;&#x3D; 3.1.
    /// </summary>
    [DataContract(Name = "ShapeOrNull")]
    public partial class ShapeOrNull : AbstractOpenAPISchema, IEquatable<ShapeOrNull>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="ShapeOrNull" /> class.
        /// </summary>
        public ShapeOrNull()
        {
            this.IsNullable = true;
            this.SchemaType= "oneOf";
        }

        /// <summary>
        /// Initializes a new instance of the <see cref="ShapeOrNull" /> class
        /// with the <see cref="Quadrilateral" /> class
        /// </summary>
        /// <param name="actualInstance">An instance of Quadrilateral.</param>
        public ShapeOrNull(Quadrilateral actualInstance)
        {
            this.IsNullable = true;
            this.SchemaType= "oneOf";
            this.ActualInstance = actualInstance;
        }

        /// <summary>
        /// Initializes a new instance of the <see cref="ShapeOrNull" /> class
        /// with the <see cref="Triangle" /> class
        /// </summary>
        /// <param name="actualInstance">An instance of Triangle.</param>
        public ShapeOrNull(Triangle actualInstance)
        {
            this.IsNullable = true;
            this.SchemaType= "oneOf";
            this.ActualInstance = actualInstance;
        }


        private Object _actualInstance;

        /// <summary>
        /// Gets or Sets ActualInstance
        /// </summary>
        public override Object ActualInstance
        {
            get
            {
                return _actualInstance;
            }
            set
            {
                if (value.GetType() == typeof(Quadrilateral))
                {
                    this._actualInstance = value;
                }
                else if (value.GetType() == typeof(Triangle))
                {
                    this._actualInstance = value;
                }
                else
                {
                    throw new ArgumentException("Invalid instance found. Must be the following types: Quadrilateral, Triangle");
                }
            }
        }

        /// <summary>
        /// Get the actual instance of `Quadrilateral`. If the actual instanct is not `Quadrilateral`,
        /// the InvalidClassException will be thrown
        /// </summary>
        /// <returns>An instance of Quadrilateral</returns>
        public Quadrilateral GetQuadrilateral()
        {
            return (Quadrilateral)this.ActualInstance;
        }

        /// <summary>
        /// Get the actual instance of `Triangle`. If the actual instanct is not `Triangle`,
        /// the InvalidClassException will be thrown
        /// </summary>
        /// <returns>An instance of Triangle</returns>
        public Triangle GetTriangle()
        {
            return (Triangle)this.ActualInstance;
        }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            var sb = new StringBuilder();
            sb.Append("class ShapeOrNull {\n");
            sb.Append("  ActualInstance: ").Append(this.ActualInstance).Append("\n");
            sb.Append("}\n");
            return sb.ToString();
        }

        /// <summary>
        /// Returns the JSON string presentation of the object
        /// </summary>
        /// <returns>JSON string presentation of the object</returns>
        public override string ToJson()
        {
            return JsonConvert.SerializeObject(this.ActualInstance, _serializerSettings);
        }

        /// <summary>
        /// Converts the JSON string into an instance of ShapeOrNull
        /// </summary>
        /// <param name="jsonString">JSON string</param>
        /// <returns>An instance of ShapeOrNull</returns>
        public static ShapeOrNull FromJson(string jsonString)
        {
            ShapeOrNull newShapeOrNull = new ShapeOrNull();

            string discriminatorValue = JObject.Parse(jsonString)["shapeType"].ToString();
            switch (discriminatorValue)
            {
                case "Quadrilateral":
                    newShapeOrNull.ActualInstance = JsonConvert.DeserializeObject<Quadrilateral>(jsonString, newShapeOrNull._serializerSettings);
                    return newShapeOrNull;
                case "Triangle":
                    newShapeOrNull.ActualInstance = JsonConvert.DeserializeObject<Triangle>(jsonString, newShapeOrNull._serializerSettings);
                    return newShapeOrNull;
                default:
                    System.Diagnostics.Debug.WriteLine(String.Format("Failed to lookup discriminator value `%s` for ShapeOrNull. Possible values: Quadrilateral Triangle", discriminatorValue));
                    break;
            }

            int match = 0;
            List<string> matchedTypes = new List<string>();

            try
            {
                newShapeOrNull.ActualInstance = JsonConvert.DeserializeObject<Quadrilateral>(jsonString, newShapeOrNull._serializerSettings);
                matchedTypes.Add("Quadrilateral");
                match++;
            }
            catch (Exception exception)
            {
                // deserialization failed, try the next one
                System.Diagnostics.Debug.WriteLine(String.Format("Failed to deserialize `%s` into Quadrilateral: %s", jsonString, exception.ToString()));
            }

            try
            {
                newShapeOrNull.ActualInstance = JsonConvert.DeserializeObject<Triangle>(jsonString, newShapeOrNull._serializerSettings);
                matchedTypes.Add("Triangle");
                match++;
            }
            catch (Exception exception)
            {
                // deserialization failed, try the next one
                System.Diagnostics.Debug.WriteLine(String.Format("Failed to deserialize `%s` into Triangle: %s", jsonString, exception.ToString()));
            }

            if (match == 0)
            {
                throw new InvalidDataException("The JSON string `" + jsonString + "` cannot be deserialized into any schema defined.");
            }
            else if (match > 1)
            {
                throw new InvalidDataException("The JSON string `" + jsonString + "` incorrectly matches more than one schema (should be exactly one match): " + matchedTypes);
            }
            
            // deserialization is considered successful at this point if no exception has been thrown.
            return newShapeOrNull;
        }

        /// <summary>
        /// Returns true if objects are equal
        /// </summary>
        /// <param name="input">Object to be compared</param>
        /// <returns>Boolean</returns>
        public override bool Equals(object input)
        {
            return OpenAPIClientUtils.compareLogic.Compare(this, input as ShapeOrNull).AreEqual;
        }

        /// <summary>
        /// Returns true if ShapeOrNull instances are equal
        /// </summary>
        /// <param name="input">Instance of ShapeOrNull to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(ShapeOrNull input)
        {
            return OpenAPIClientUtils.compareLogic.Compare(this, input).AreEqual;
        }

        /// <summary>
        /// Gets the hash code
        /// </summary>
        /// <returns>Hash code</returns>
        public override int GetHashCode()
        {
            unchecked // Overflow is fine, just wrap
            {
                int hashCode = 41;
                if (this.ActualInstance != null)
                    hashCode = hashCode * 59 + this.ActualInstance.GetHashCode();
                return hashCode;
            }
        }

        /// <summary>
        /// To validate all properties of the instance
        /// </summary>
        /// <param name="validationContext">Validation context</param>
        /// <returns>Validation Result</returns>
        IEnumerable<System.ComponentModel.DataAnnotations.ValidationResult> IValidatableObject.Validate(ValidationContext validationContext)
        {
            yield break;
        }
    }

}
