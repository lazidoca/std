import { Link, LinkProps } from "@chakra-ui/next-js";
import { IconButton, IconButtonProps, forwardRef } from "@chakra-ui/react";
import {
  ArrowLeftIcon,
  BookmarkIcon,
  CloudArrowUpIcon,
  EllipsisHorizontalIcon,
  PencilIcon,
  PlusIcon,
  XMarkIcon,
} from "@heroicons/react/24/outline";
import { BookmarkIcon as BookmarkSolidIcon } from "@heroicons/react/24/solid";
import { MouseEvent, MouseEventHandler, useCallback } from "react";

import { SendIcon } from "../../graphics/SendIcon";

const actionStyles = {
  width: 8,
  height: 8,
  display: "flex",
  justifyContent: "center",
  alignItems: "center",
  borderRadius: "full",
  p: 1,
  _hover: { bgColor: "blackAlpha.50" },
};

function useClickHandler(onClick: MouseEventHandler | undefined) {
  // This allows us to progressively enhance features on the application by
  // treating important buttons as links to fallback pages. For example, there
  // may be a button that triggers the opening of a modal dialogue but if the
  // user has JavaScript disabled due to device constraints or privacy reasons,
  // the functionality must also be implemented by a normal page.
  return useCallback(
    (e: MouseEvent) => {
      if (onClick) {
        e.preventDefault();
        return onClick?.(e);
      }
    },
    [onClick],
  );
}

export const Action = forwardRef(
  ({ children, onClick, ...props }: LinkProps, ref) => {
    const handleClick = useClickHandler(onClick);
    return (
      <Link ref={ref} onClick={handleClick} {...actionStyles} {...props}>
        {children}
      </Link>
    );
  },
);

export const ActionButton = forwardRef<IconButtonProps, "button">(
  ({ children, ...props }, ref) => {
    return (
      <IconButton
        ref={ref}
        width={8}
        height={8}
        display="flex"
        justifyContent="center"
        alignItems="center"
        borderRadius="full"
        p={1}
        bgColor="transparent"
        _hover={{ bgColor: "blackAlpha.50" }}
        {...props}
      >
        {children}
      </IconButton>
    );
  },
);

// A few actions have default page destinations (partly for consistency and also
// for accessibility and no-JS modes) so this just redefines `href` as optional.
type WithOptionalURL = Omit<LinkProps, "href"> & {
  href?: string | undefined;
};

export type WithOptionalARIALabel = Omit<IconButtonProps, "aria-label"> & {
  "aria-label"?: string | undefined;
};

export const Edit = forwardRef(
  ({ "aria-label": al, ...props }: WithOptionalARIALabel, ref) => {
    return (
      <ActionButton ref={ref} title="Edit" aria-label={al ?? "edit"} {...props}>
        <PencilIcon width="1em" height="1em" />
      </ActionButton>
    );
  },
);

export function Close({ "aria-label": al, ...props }: WithOptionalARIALabel) {
  return (
    <ActionButton
      size="sm"
      title="Close"
      aria-label={al ?? "close"}
      {...props}
      icon={<XMarkIcon width="1.4em" />}
    />
  );
}

export function Back({
  href,
  "aria-label": al,
  ...props
}: WithOptionalARIALabel & WithOptionalURL) {
  if (href)
    return (
      <Action
        href={href}
        size="sm"
        title="Back"
        aria-label={al ?? "back"}
        {...props}
      >
        <ArrowLeftIcon width="1.4em" />
      </Action>
    );

  return (
    <ActionButton
      size="sm"
      title="Back"
      aria-label={al ?? "back"}
      {...props}
      icon={<ArrowLeftIcon width="1.4em" />}
    />
  );
}

export function Send({ "aria-label": al, ...props }: WithOptionalARIALabel) {
  return (
    <ActionButton
      size="sm"
      title="Send"
      aria-label={al ?? "send"}
      {...props}
      icon={<SendIcon width="1.4em" />}
    />
  );
}

export function Save({ "aria-label": al, ...props }: WithOptionalARIALabel) {
  return (
    <ActionButton
      size="sm"
      title="Save"
      aria-label={al ?? "save"}
      {...props}
      icon={<CloudArrowUpIcon width="1.4em" />}
    />
  );
}

export function Add({ "aria-label": al, ...props }: WithOptionalARIALabel) {
  return (
    <ActionButton
      size="sm"
      title="Add"
      aria-label={al ?? "add"}
      {...props}
      icon={<PlusIcon width="1.4em" />}
    />
  );
}

// NOTE: this one is forward-ref'd because it's used as a chakra Menu button.
// https://chakra-ui.com/docs/components/menu#customizing-the-button
export const More = forwardRef(
  ({ "aria-label": al, ...props }: WithOptionalARIALabel, ref) => {
    return (
      <ActionButton
        ref={ref}
        size="sm"
        title="More"
        aria-label={al ?? "more"}
        {...props}
        icon={<EllipsisHorizontalIcon width="1.4em" />}
      />
    );
  },
);

export type BookmarkProps = {
  bookmarked: boolean;
} & WithOptionalARIALabel;

export const Bookmark = forwardRef(
  ({ "aria-label": al, bookmarked, ...props }: BookmarkProps, ref) => {
    return (
      <ActionButton
        ref={ref}
        size="sm"
        title={bookmarked ? "In collection" : "Save to collection"}
        aria-label={al ?? "save"}
        {...props}
        icon={
          bookmarked ? (
            <BookmarkSolidIcon width="1.4em" />
          ) : (
            <BookmarkIcon width="1.4em" />
          )
        }
      />
    );
  },
);
