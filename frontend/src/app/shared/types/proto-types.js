/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
"use strict";

var $protobuf = require("protobufjs/minimal");

// Common aliases
var $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
var $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

$root.models = (function() {

    /**
     * Namespace models.
     * @exports models
     * @namespace
     */
    var models = {};

    models.User = (function() {

        /**
         * Properties of a User.
         * @memberof models
         * @interface IUser
         * @property {number|null} [ID] User ID
         * @property {number|null} [CreatedAt] User CreatedAt
         * @property {number|null} [UpdatedAt] User UpdatedAt
         * @property {string|null} [FirstName] User FirstName
         * @property {string|null} [LastName] User LastName
         * @property {string|null} [Username] User Username
         * @property {string|null} [Password] User Password
         * @property {models.User.AuthRole|null} [Role] User Role
         */

        /**
         * Constructs a new User.
         * @memberof models
         * @classdesc Represents a User.
         * @implements IUser
         * @constructor
         * @param {models.IUser=} [properties] Properties to set
         */
        function User(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * User ID.
         * @member {number} ID
         * @memberof models.User
         * @instance
         */
        User.prototype.ID = 0;

        /**
         * User CreatedAt.
         * @member {number} CreatedAt
         * @memberof models.User
         * @instance
         */
        User.prototype.CreatedAt = 0;

        /**
         * User UpdatedAt.
         * @member {number} UpdatedAt
         * @memberof models.User
         * @instance
         */
        User.prototype.UpdatedAt = 0;

        /**
         * User FirstName.
         * @member {string} FirstName
         * @memberof models.User
         * @instance
         */
        User.prototype.FirstName = "";

        /**
         * User LastName.
         * @member {string} LastName
         * @memberof models.User
         * @instance
         */
        User.prototype.LastName = "";

        /**
         * User Username.
         * @member {string} Username
         * @memberof models.User
         * @instance
         */
        User.prototype.Username = "";

        /**
         * User Password.
         * @member {string} Password
         * @memberof models.User
         * @instance
         */
        User.prototype.Password = "";

        /**
         * User Role.
         * @member {models.User.AuthRole} Role
         * @memberof models.User
         * @instance
         */
        User.prototype.Role = 0;

        /**
         * Creates a new User instance using the specified properties.
         * @function create
         * @memberof models.User
         * @static
         * @param {models.IUser=} [properties] Properties to set
         * @returns {models.User} User instance
         */
        User.create = function create(properties) {
            return new User(properties);
        };

        /**
         * Encodes the specified User message. Does not implicitly {@link models.User.verify|verify} messages.
         * @function encode
         * @memberof models.User
         * @static
         * @param {models.IUser} message User message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        User.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.ID != null && message.hasOwnProperty("ID"))
                writer.uint32(/* id 1, wireType 0 =*/8).uint32(message.ID);
            if (message.CreatedAt != null && message.hasOwnProperty("CreatedAt"))
                writer.uint32(/* id 2, wireType 0 =*/16).uint32(message.CreatedAt);
            if (message.UpdatedAt != null && message.hasOwnProperty("UpdatedAt"))
                writer.uint32(/* id 3, wireType 0 =*/24).uint32(message.UpdatedAt);
            if (message.FirstName != null && message.hasOwnProperty("FirstName"))
                writer.uint32(/* id 4, wireType 2 =*/34).string(message.FirstName);
            if (message.LastName != null && message.hasOwnProperty("LastName"))
                writer.uint32(/* id 5, wireType 2 =*/42).string(message.LastName);
            if (message.Username != null && message.hasOwnProperty("Username"))
                writer.uint32(/* id 6, wireType 2 =*/50).string(message.Username);
            if (message.Password != null && message.hasOwnProperty("Password"))
                writer.uint32(/* id 7, wireType 2 =*/58).string(message.Password);
            if (message.Role != null && message.hasOwnProperty("Role"))
                writer.uint32(/* id 8, wireType 0 =*/64).int32(message.Role);
            return writer;
        };

        /**
         * Encodes the specified User message, length delimited. Does not implicitly {@link models.User.verify|verify} messages.
         * @function encodeDelimited
         * @memberof models.User
         * @static
         * @param {models.IUser} message User message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        User.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a User message from the specified reader or buffer.
         * @function decode
         * @memberof models.User
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {models.User} User
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        User.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.models.User();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.ID = reader.uint32();
                    break;
                case 2:
                    message.CreatedAt = reader.uint32();
                    break;
                case 3:
                    message.UpdatedAt = reader.uint32();
                    break;
                case 4:
                    message.FirstName = reader.string();
                    break;
                case 5:
                    message.LastName = reader.string();
                    break;
                case 6:
                    message.Username = reader.string();
                    break;
                case 7:
                    message.Password = reader.string();
                    break;
                case 8:
                    message.Role = reader.int32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a User message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof models.User
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {models.User} User
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        User.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a User message.
         * @function verify
         * @memberof models.User
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        User.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.ID != null && message.hasOwnProperty("ID"))
                if (!$util.isInteger(message.ID))
                    return "ID: integer expected";
            if (message.CreatedAt != null && message.hasOwnProperty("CreatedAt"))
                if (!$util.isInteger(message.CreatedAt))
                    return "CreatedAt: integer expected";
            if (message.UpdatedAt != null && message.hasOwnProperty("UpdatedAt"))
                if (!$util.isInteger(message.UpdatedAt))
                    return "UpdatedAt: integer expected";
            if (message.FirstName != null && message.hasOwnProperty("FirstName"))
                if (!$util.isString(message.FirstName))
                    return "FirstName: string expected";
            if (message.LastName != null && message.hasOwnProperty("LastName"))
                if (!$util.isString(message.LastName))
                    return "LastName: string expected";
            if (message.Username != null && message.hasOwnProperty("Username"))
                if (!$util.isString(message.Username))
                    return "Username: string expected";
            if (message.Password != null && message.hasOwnProperty("Password"))
                if (!$util.isString(message.Password))
                    return "Password: string expected";
            if (message.Role != null && message.hasOwnProperty("Role"))
                switch (message.Role) {
                default:
                    return "Role: enum value expected";
                case 0:
                case 1:
                case 2:
                    break;
                }
            return null;
        };

        /**
         * Creates a User message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof models.User
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {models.User} User
         */
        User.fromObject = function fromObject(object) {
            if (object instanceof $root.models.User)
                return object;
            var message = new $root.models.User();
            if (object.ID != null)
                message.ID = object.ID >>> 0;
            if (object.CreatedAt != null)
                message.CreatedAt = object.CreatedAt >>> 0;
            if (object.UpdatedAt != null)
                message.UpdatedAt = object.UpdatedAt >>> 0;
            if (object.FirstName != null)
                message.FirstName = String(object.FirstName);
            if (object.LastName != null)
                message.LastName = String(object.LastName);
            if (object.Username != null)
                message.Username = String(object.Username);
            if (object.Password != null)
                message.Password = String(object.Password);
            switch (object.Role) {
            case "ROLE_USER":
            case 0:
                message.Role = 0;
                break;
            case "ROLE_CONTENT_ADMIN":
            case 1:
                message.Role = 1;
                break;
            case "ROLE_ADMIN":
            case 2:
                message.Role = 2;
                break;
            }
            return message;
        };

        /**
         * Creates a plain object from a User message. Also converts values to other types if specified.
         * @function toObject
         * @memberof models.User
         * @static
         * @param {models.User} message User
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        User.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.ID = 0;
                object.CreatedAt = 0;
                object.UpdatedAt = 0;
                object.FirstName = "";
                object.LastName = "";
                object.Username = "";
                object.Password = "";
                object.Role = options.enums === String ? "ROLE_USER" : 0;
            }
            if (message.ID != null && message.hasOwnProperty("ID"))
                object.ID = message.ID;
            if (message.CreatedAt != null && message.hasOwnProperty("CreatedAt"))
                object.CreatedAt = message.CreatedAt;
            if (message.UpdatedAt != null && message.hasOwnProperty("UpdatedAt"))
                object.UpdatedAt = message.UpdatedAt;
            if (message.FirstName != null && message.hasOwnProperty("FirstName"))
                object.FirstName = message.FirstName;
            if (message.LastName != null && message.hasOwnProperty("LastName"))
                object.LastName = message.LastName;
            if (message.Username != null && message.hasOwnProperty("Username"))
                object.Username = message.Username;
            if (message.Password != null && message.hasOwnProperty("Password"))
                object.Password = message.Password;
            if (message.Role != null && message.hasOwnProperty("Role"))
                object.Role = options.enums === String ? $root.models.User.AuthRole[message.Role] : message.Role;
            return object;
        };

        /**
         * Converts this User to JSON.
         * @function toJSON
         * @memberof models.User
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        User.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * AuthRole enum.
         * @name models.User.AuthRole
         * @enum {string}
         * @property {number} ROLE_USER=0 ROLE_USER value
         * @property {number} ROLE_CONTENT_ADMIN=1 ROLE_CONTENT_ADMIN value
         * @property {number} ROLE_ADMIN=2 ROLE_ADMIN value
         */
        User.AuthRole = (function() {
            var valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "ROLE_USER"] = 0;
            values[valuesById[1] = "ROLE_CONTENT_ADMIN"] = 1;
            values[valuesById[2] = "ROLE_ADMIN"] = 2;
            return values;
        })();

        return User;
    })();

    return models;
})();

module.exports = $root;
