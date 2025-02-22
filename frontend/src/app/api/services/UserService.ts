/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { care_vault_dtos_user_dtos_UserModifyDto } from '../models/care_vault_dtos_user_dtos_UserModifyDto';
import type { care_vault_dtos_user_dtos_UserPublicDto } from '../models/care_vault_dtos_user_dtos_UserPublicDto';
import type { models_User } from '../models/models_User';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class UserService {
    /**
     * Create a new user
     * Creates a new user with the provided name and email. Ensures the email is unique.
     * @param user User data
     * @returns any User created successfully
     * @throws ApiError
     */
    public static postUser(
        user: care_vault_dtos_user_dtos_UserPublicDto,
    ): CancelablePromise<any> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/user',
            body: user,
            errors: {
                400: `Invalid input or email already in use`,
                500: `Could not create user`,
            },
        });
    }
    /**
     * Modify a user.
     * Modify user user data by providing new data and the user-ID.
     * @param user User data
     * @returns models_User User successfully modified.
     * @throws ApiError
     */
    public static patchUser(
        user: care_vault_dtos_user_dtos_UserModifyDto,
    ): CancelablePromise<models_User> {
        return __request(OpenAPI, {
            method: 'PATCH',
            url: '/user',
            body: user,
            errors: {
                400: `Invalid ID format`,
                404: `Database query failed`,
            },
        });
    }
    /**
     * Get a user by ID
     * Get detailed information about a user
     * @param id User ID
     * @returns models_User Successfully retrieved user
     * @throws ApiError
     */
    public static getUser(
        id: number,
    ): CancelablePromise<models_User> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/user/{id}',
            path: {
                'id': id,
            },
            errors: {
                400: `Invalid ID format`,
                404: `Database query failed`,
            },
        });
    }
}
