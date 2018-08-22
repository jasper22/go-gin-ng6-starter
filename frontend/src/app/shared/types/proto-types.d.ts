import * as $protobuf from "protobufjs";
/** Namespace models. */
export namespace models {

    /** Properties of a User. */
    interface IUser {

        /** User ID */
        ID?: (number|null);

        /** User CreatedAt */
        CreatedAt?: (number|null);

        /** User UpdatedAt */
        UpdatedAt?: (number|null);

        /** User FirstName */
        FirstName?: (string|null);

        /** User LastName */
        LastName?: (string|null);

        /** User Username */
        Username?: (string|null);

        /** User Password */
        Password?: (string|null);

        /** User Role */
        Role?: (models.User.AuthRole|null);
    }

    /** Represents a User. */
    class User implements IUser {

        /**
         * Constructs a new User.
         * @param [properties] Properties to set
         */
        constructor(properties?: models.IUser);

        /** User ID. */
        public ID: number;

        /** User CreatedAt. */
        public CreatedAt: number;

        /** User UpdatedAt. */
        public UpdatedAt: number;

        /** User FirstName. */
        public FirstName: string;

        /** User LastName. */
        public LastName: string;

        /** User Username. */
        public Username: string;

        /** User Password. */
        public Password: string;

        /** User Role. */
        public Role: models.User.AuthRole;

        /**
         * Creates a new User instance using the specified properties.
         * @param [properties] Properties to set
         * @returns User instance
         */
        public static create(properties?: models.IUser): models.User;

        /**
         * Encodes the specified User message. Does not implicitly {@link models.User.verify|verify} messages.
         * @param message User message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: models.IUser, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified User message, length delimited. Does not implicitly {@link models.User.verify|verify} messages.
         * @param message User message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: models.IUser, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a User message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns User
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): models.User;

        /**
         * Decodes a User message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns User
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): models.User;

        /**
         * Verifies a User message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a User message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns User
         */
        public static fromObject(object: { [k: string]: any }): models.User;

        /**
         * Creates a plain object from a User message. Also converts values to other types if specified.
         * @param message User
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: models.User, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this User to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }

    namespace User {

        /** AuthRole enum. */
        enum AuthRole {
            ROLE_USER = 0,
            ROLE_CONTENT_ADMIN = 1,
            ROLE_ADMIN = 2
        }
    }
}
