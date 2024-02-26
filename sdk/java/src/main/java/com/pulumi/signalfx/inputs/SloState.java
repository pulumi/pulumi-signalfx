// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.signalfx.inputs.SloInputArgs;
import com.pulumi.signalfx.inputs.SloTargetArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SloState extends com.pulumi.resources.ResourceArgs {

    public static final SloState Empty = new SloState();

    /**
     * Description of the SLO
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the SLO
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * SignalFlow program and arguments text strings that define the streams used as successful event count and total event
     * count
     * 
     */
    @Import(name="input")
    private @Nullable Output<SloInputArgs> input;

    /**
     * @return SignalFlow program and arguments text strings that define the streams used as successful event count and total event
     * count
     * 
     */
    public Optional<Output<SloInputArgs>> input() {
        return Optional.ofNullable(this.input);
    }

    /**
     * Name of the SLO
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the SLO
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Define target value of the service level indicator in the appropriate time period.
     * 
     */
    @Import(name="target")
    private @Nullable Output<SloTargetArgs> target;

    /**
     * @return Define target value of the service level indicator in the appropriate time period.
     * 
     */
    public Optional<Output<SloTargetArgs>> target() {
        return Optional.ofNullable(this.target);
    }

    /**
     * Type of the SLO. Currently only RequestBased SLO is supported
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return Type of the SLO. Currently only RequestBased SLO is supported
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private SloState() {}

    private SloState(SloState $) {
        this.description = $.description;
        this.input = $.input;
        this.name = $.name;
        this.target = $.target;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SloState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SloState $;

        public Builder() {
            $ = new SloState();
        }

        public Builder(SloState defaults) {
            $ = new SloState(Objects.requireNonNull(defaults));
        }

        /**
         * @param description Description of the SLO
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the SLO
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param input SignalFlow program and arguments text strings that define the streams used as successful event count and total event
         * count
         * 
         * @return builder
         * 
         */
        public Builder input(@Nullable Output<SloInputArgs> input) {
            $.input = input;
            return this;
        }

        /**
         * @param input SignalFlow program and arguments text strings that define the streams used as successful event count and total event
         * count
         * 
         * @return builder
         * 
         */
        public Builder input(SloInputArgs input) {
            return input(Output.of(input));
        }

        /**
         * @param name Name of the SLO
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the SLO
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param target Define target value of the service level indicator in the appropriate time period.
         * 
         * @return builder
         * 
         */
        public Builder target(@Nullable Output<SloTargetArgs> target) {
            $.target = target;
            return this;
        }

        /**
         * @param target Define target value of the service level indicator in the appropriate time period.
         * 
         * @return builder
         * 
         */
        public Builder target(SloTargetArgs target) {
            return target(Output.of(target));
        }

        /**
         * @param type Type of the SLO. Currently only RequestBased SLO is supported
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Type of the SLO. Currently only RequestBased SLO is supported
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public SloState build() {
            return $;
        }
    }

}
