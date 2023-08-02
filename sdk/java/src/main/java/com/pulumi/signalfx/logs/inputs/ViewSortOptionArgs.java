// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.logs.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;


public final class ViewSortOptionArgs extends com.pulumi.resources.ResourceArgs {

    public static final ViewSortOptionArgs Empty = new ViewSortOptionArgs();

    @Import(name="descending", required=true)
    private Output<Boolean> descending;

    public Output<Boolean> descending() {
        return this.descending;
    }

    @Import(name="field", required=true)
    private Output<String> field;

    public Output<String> field() {
        return this.field;
    }

    private ViewSortOptionArgs() {}

    private ViewSortOptionArgs(ViewSortOptionArgs $) {
        this.descending = $.descending;
        this.field = $.field;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ViewSortOptionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ViewSortOptionArgs $;

        public Builder() {
            $ = new ViewSortOptionArgs();
        }

        public Builder(ViewSortOptionArgs defaults) {
            $ = new ViewSortOptionArgs(Objects.requireNonNull(defaults));
        }

        public Builder descending(Output<Boolean> descending) {
            $.descending = descending;
            return this;
        }

        public Builder descending(Boolean descending) {
            return descending(Output.of(descending));
        }

        public Builder field(Output<String> field) {
            $.field = field;
            return this;
        }

        public Builder field(String field) {
            return field(Output.of(field));
        }

        public ViewSortOptionArgs build() {
            $.descending = Objects.requireNonNull($.descending, "expected parameter 'descending' to be non-null");
            $.field = Objects.requireNonNull($.field, "expected parameter 'field' to be non-null");
            return $;
        }
    }

}
