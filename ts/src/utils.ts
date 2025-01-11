/**
 * Narrow the type of the provided argument
 * to T in the caller's scope.
 *
 * Calling this function does nothing at runtime, it is only useful for the typescript compiler.
 */
export function narrow<T>(_: unknown): asserts _ is T {}
