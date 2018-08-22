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
         * @property {number|null} [createdAt] User createdAt
         * @property {number|null} [updatedAt] User updatedAt
         * @property {string|null} [firstName] User firstName
         * @property {string|null} [lastName] User lastName
         * @property {string|null} [username] User username
         * @property {string|null} [password] User password
         * @property {models.User.Role|null} [role] User role
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
         * User createdAt.
         * @member {number} createdAt
         * @memberof models.User
         * @instance
         */
        User.prototype.createdAt = 0;

        /**
         * User updatedAt.
         * @member {number} updatedAt
         * @memberof models.User
         * @instance
         */
        User.prototype.updatedAt = 0;

        /**
         * User firstName.
         * @member {string} firstName
         * @memberof models.User
         * @instance
         */
        User.prototype.firstName = "";

        /**
         * User lastName.
         * @member {string} lastName
         * @memberof models.User
         * @instance
         */
        User.prototype.lastName = "";

        /**
         * User username.
         * @member {string} username
         * @memberof models.User
         * @instance
         */
        User.prototype.username = "";

        /**
         * User password.
         * @member {string} password
         * @memberof models.User
         * @instance
         */
        User.prototype.password = "";

        /**
         * User role.
         * @member {models.User.Role} role
         * @memberof models.User
         * @instance
         */
        User.prototype.role = 0;

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
            if (message.createdAt != null && message.hasOwnProperty("createdAt"))
                writer.uint32(/* id 2, wireType 0 =*/16).uint32(message.createdAt);
            if (message.updatedAt != null && message.hasOwnProperty("updatedAt"))
                writer.uint32(/* id 3, wireType 0 =*/24).uint32(message.updatedAt);
            if (message.firstName != null && message.hasOwnProperty("firstName"))
                writer.uint32(/* id 4, wireType 2 =*/34).string(message.firstName);
            if (message.lastName != null && message.hasOwnProperty("lastName"))
                writer.uint32(/* id 5, wireType 2 =*/42).string(message.lastName);
            if (message.username != null && message.hasOwnProperty("username"))
                writer.uint32(/* id 6, wireType 2 =*/50).string(message.username);
            if (message.password != null && message.hasOwnProperty("password"))
                writer.uint32(/* id 7, wireType 2 =*/58).string(message.password);
            if (message.role != null && message.hasOwnProperty("role"))
                writer.uint32(/* id 8, wireType 0 =*/64).int32(message.role);
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
                    message.createdAt = reader.uint32();
                    break;
                case 3:
                    message.updatedAt = reader.uint32();
                    break;
                case 4:
                    message.firstName = reader.string();
                    break;
                case 5:
                    message.lastName = reader.string();
                    break;
                case 6:
                    message.username = reader.string();
                    break;
                case 7:
                    message.password = reader.string();
                    break;
                case 8:
                    message.role = reader.int32();
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
            if (message.createdAt != null && message.hasOwnProperty("createdAt"))
                if (!$util.isInteger(message.createdAt))
                    return "createdAt: integer expected";
            if (message.updatedAt != null && message.hasOwnProperty("updatedAt"))
                if (!$util.isInteger(message.updatedAt))
                    return "updatedAt: integer expected";
            if (message.firstName != null && message.hasOwnProperty("firstName"))
                if (!$util.isString(message.firstName))
                    return "firstName: string expected";
            if (message.lastName != null && message.hasOwnProperty("lastName"))
                if (!$util.isString(message.lastName))
                    return "lastName: string expected";
            if (message.username != null && message.hasOwnProperty("username"))
                if (!$util.isString(message.username))
                    return "username: string expected";
            if (message.password != null && message.hasOwnProperty("password"))
                if (!$util.isString(message.password))
                    return "password: string expected";
            if (message.role != null && message.hasOwnProperty("role"))
                switch (message.role) {
                default:
                    return "role: enum value expected";
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
            if (object.createdAt != null)
                message.createdAt = object.createdAt >>> 0;
            if (object.updatedAt != null)
                message.updatedAt = object.updatedAt >>> 0;
            if (object.firstName != null)
                message.firstName = String(object.firstName);
            if (object.lastName != null)
                message.lastName = String(object.lastName);
            if (object.username != null)
                message.username = String(object.username);
            if (object.password != null)
                message.password = String(object.password);
            switch (object.role) {
            case "ROLE_USER":
            case 0:
                message.role = 0;
                break;
            case "ROLE_CONTENT_ADMIN":
            case 1:
                message.role = 1;
                break;
            case "ROLE_ADMIN":
            case 2:
                message.role = 2;
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
                object.createdAt = 0;
                object.updatedAt = 0;
                object.firstName = "";
                object.lastName = "";
                object.username = "";
                object.password = "";
                object.role = options.enums === String ? "ROLE_USER" : 0;
            }
            if (message.ID != null && message.hasOwnProperty("ID"))
                object.ID = message.ID;
            if (message.createdAt != null && message.hasOwnProperty("createdAt"))
                object.createdAt = message.createdAt;
            if (message.updatedAt != null && message.hasOwnProperty("updatedAt"))
                object.updatedAt = message.updatedAt;
            if (message.firstName != null && message.hasOwnProperty("firstName"))
                object.firstName = message.firstName;
            if (message.lastName != null && message.hasOwnProperty("lastName"))
                object.lastName = message.lastName;
            if (message.username != null && message.hasOwnProperty("username"))
                object.username = message.username;
            if (message.password != null && message.hasOwnProperty("password"))
                object.password = message.password;
            if (message.role != null && message.hasOwnProperty("role"))
                object.role = options.enums === String ? $root.models.User.Role[message.role] : message.role;
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
         * Role enum.
         * @name models.User.Role
         * @enum {string}
         * @property {number} ROLE_USER=0 ROLE_USER value
         * @property {number} ROLE_CONTENT_ADMIN=1 ROLE_CONTENT_ADMIN value
         * @property {number} ROLE_ADMIN=2 ROLE_ADMIN value
         */
        User.Role = (function() {
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
